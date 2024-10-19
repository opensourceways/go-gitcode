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
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.gitcode.com/api/v5/"

	HeaderAuthorization        = "Authorization"
	HeaderUserAgentName        = "User-Agent"
	HeaderUserAgentValue       = "OpenSourceCommunityRobot/1.0.0"
	HeaderMediaTypeName        = "Accept"
	HeaderMediaTypeValue       = "application/json"
	HeaderContentTypeName      = "Content-Type"
	HeaderContentTypeJsonValue = "application/json"
	HeaderContentTypeFormValue = "application/x-www-form-urlencoded"
)

var (
	nilContentError    = errors.New("request context should be non-nil")
	pathForbiddenError = errors.New("post http request body should be non-nil")
)

type RequestHandlerType string

const (
	Json  RequestHandlerType = "json"
	Form  RequestHandlerType = "form"
	Query RequestHandlerType = "query"
)

type RequestHandler struct {
	buf io.ReadWriter
	t   RequestHandlerType
}

func (b RequestHandler) PreOperate(uri *url.URL, body any) error {
	var err error
	if body != nil {
		switch b.t {
		case Query:
			// set url query
			if body, ok := body.(*url.Values); ok {
				uri.RawQuery = body.Encode()
			}
		case Form:
			// set form data
			if body, ok := body.(*url.Values); ok {
				b.buf = bytes.NewBufferString(body.Encode())
			}
		default:
			// set json data
			b.buf = &bytes.Buffer{}
			enc := json.NewEncoder(b.buf)
			enc.SetEscapeHTML(false)
			err = enc.Encode(body)
		}
	}
	return err
}

func (b RequestHandler) PostOperate(req *http.Request) {
	switch b.t {
	case Form:
		req.Header.Set(HeaderContentTypeName, HeaderContentTypeFormValue)
	case Query:
		fallthrough
	default:
		req.Header.Set(HeaderContentTypeName, HeaderContentTypeJsonValue)
	}
}

type service struct {
	api *APIClient
}

type IssuesService service

type PullRequestsService service

type RepositoryService service
