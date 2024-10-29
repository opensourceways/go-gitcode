// Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package openapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"
)

// APIClient 一个客户端，管理与 GitCode OpenAPI 的通信
type APIClient struct {
	client  *http.Client
	baseURL *url.URL

	common service // Reuse a single struct instead of allocating one for each service on the heap.
	// 各个模块的服务
	Issues       *IssuesService
	PullRequests *PullRequestsService
	Repository   *RepositoryService
	User         *UserService
}

// roundTripperFunc creates a RoundTripper (transport)
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (fn roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return fn(r)
}

func NewAPIClientWithAuthorization(token []byte) *APIClient {
	httpClient := &http.Client{
		Transport: roundTripperFunc(
			func(req *http.Request) (*http.Response, error) {
				req = req.Clone(req.Context())
				req.Header.Set(headerAuthorization, "Bearer "+string(token))
				req.Header.Set(headerUserAgentName, headerUserAgentValue)
				req.Header.Set(headerMediaTypeName, headerMediaTypeValue)
				return createTransport(nil).RoundTrip(req)
			},
		),
		Timeout: 90 * time.Second,
	}

	c := &APIClient{client: httpClient}

	if c.baseURL == nil {
		c.baseURL, _ = url.Parse(defaultBaseURL)
	}
	c.common.api = c

	c.Issues = (*IssuesService)(&c.common)
	c.PullRequests = (*PullRequestsService)(&c.common)
	c.Repository = (*RepositoryService)(&c.common)
	c.User = (*UserService)(&c.common)

	return c
}

func createTransport(localAddr net.Addr) *http.Transport {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	if localAddr != nil {
		dialer.LocalAddr = localAddr
	}
	return &http.Transport{
		DialContext:           transportDialContext(dialer),
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

func transportDialContext(dialer *net.Dialer) func(context.Context, string, string) (net.Conn, error) {
	return dialer.DialContext
}

func newRequest(c *APIClient, method, urlStr string, body any, handlers ...RequestHandler) (*http.Request, error) {
	uri, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var handler *RequestHandler
	// 如果 body 不为空，默认以 JSON 处理
	if len(handlers) == 0 {
		handler = &RequestHandler{t: Json}
	} else {
		handler = &handlers[0]
	}
	// 处理 url query、表单、json
	if err = handler.PreOperate(uri, body); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, uri.String(), handler.buf)
	if err != nil {
		return nil, err
	}

	// 添加 header
	handler.PostOperate(req)
	return req, nil
}

func (c *APIClient) Do(ctx context.Context, req *http.Request, receiver any) (*http.Response, error) {
	var resp *http.Response
	var err error

	retry := 3
	for i := 1; i <= retry; i++ {
		resp, err = c.BareDo(ctx, req)
		if resp != nil && resp.StatusCode <= http.StatusUnavailableForLegalReasons {
			break
		}
		time.Sleep(time.Duration(i) * time.Second)
	}
	if err == nil && resp != nil && resp.StatusCode >= http.StatusMultipleChoices {
		var str strings.Builder
		_, _ = io.Copy(&str, resp.Body)
		err = errors.New(str.String())
		_ = resp.Body.Close()
	}
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	switch receiver := receiver.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(receiver, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(receiver)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			fmt.Println(decErr)
			err = decErr
		}
	}
	return resp, err
}

/*
下表显示了 GitCode api 请求可能的返回代码：https://docs.gitcode.com/docs/openapi/guide/
返回值	描述
200 OK	GET、PUT 或 DELETE 请求成功，并且资源本身以 JSON 形式返回
201 Created	POST 请求成功，并且资源以 JSON 形式返回
202 Accepted	GET、PUT 或 DELETE 请求成功，并且资源计划进行处理
204 No Content	服务器已成功满足请求，并且在响应负载体中没有额外的内容发送
301 Moved Permanently	资源已被定位到由 Location 头给出的 URL
304 Not Modified	资源自上次请求以来未被修改
400 Bad Request	api 请求的必需属性缺失。例如，未给出问题的标题
401 Unauthorized	用户未经认证。需要有效的用户令牌
403 Forbidden	请求不被允许。例如，用户不允许删除项目
404 Not Found	无法访问资源。例如，无法找到资源的 ID，或者用户无权访问资源
405 Method Not Allowed	不支持请求
409 Conflict	冲突的资源已存在。例如，创建已存在名称的项目
412 Precondition Failed	请求被拒绝。这可能发生在尝试删除在此期间被修改的资源时，如果提供了 If-Unmodified-Since 头
422 Unprocessable	无法处理实体
429 Too Many Requests	用户超过了应用程序速率限制
500 Server Error	处理请求时，服务器上出了问题
503 Service Unavailable	服务器无法处理请求，因为服务器暂时过载
*/
func successGetData(resp *http.Response) bool {
	return resp != nil && resp.StatusCode == http.StatusOK
}

func successCreated(resp *http.Response) bool {
	return resp != nil && resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated
}

func successModified(resp *http.Response) bool {
	return resp != nil && resp.StatusCode == http.StatusOK ||
		resp.StatusCode == http.StatusAccepted || resp.StatusCode == http.StatusNoContent
}

func (c *APIClient) BareDo(ctx context.Context, req *http.Request) (*http.Response, error) {
	if ctx == nil {
		return nil, nilContentError
	}
	req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		// canceled or times out, ctx.Err() will be returned
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	return resp, err
}
