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
	"reflect"
)

const (
	defaultBaseURL = "https://api.gitcode.com/api/v5/"

	headerAuthorization        = "Authorization"
	headerUserAgentName        = "User-Agent"
	headerUserAgentValue       = "OpenSourceCommunityRobot/1.0.0"
	headerMediaTypeName        = "Accept"
	headerMediaTypeValue       = "application/json"
	headerContentTypeName      = "Content-Type"
	headerContentTypeJsonValue = "application/json"
	headerContentTypeFormValue = "application/x-www-form-urlencoded"
)

var (
	nilContentError               = errors.New("request context should be non-nil")
	respReceiverNotAnPointerError = errors.New("response's receiver should be an pointer")
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

func (b *RequestHandler) PreOperate(uri *url.URL, body any) error {
	var err error
	if body != nil {

		if reflect.TypeOf(body).Kind() != reflect.Pointer {
			return respReceiverNotAnPointerError
		}

		switch b.t {
		case Query:
			// set url query
			if query, ok := body.(*url.Values); ok {
				uri.RawQuery = query.Encode()
			}
		case Form:
			// set form data
			b.buf = buildRequestForm(body)
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

func buildRequestForm(body any) *bytes.Buffer {
	if body == nil {
		return nil
	}
	form, ok := body.(*url.Values)
	if !ok {
		v := reflect.ValueOf(body)
		mt := v.MethodByName("Form")
		if mt.IsValid() {
			vs := mt.Call(nil)
			p := vs[0].Interface()
			form, ok = p.(*url.Values)
		}
	}
	if ok {
		return bytes.NewBufferString(form.Encode())
	}

	return nil
}

func (b *RequestHandler) PostOperate(req *http.Request) {
	switch b.t {
	case Form:
		req.Header.Set(headerContentTypeName, headerContentTypeFormValue)
	case Query:
		fallthrough
	default:
		req.Header.Set(headerContentTypeName, headerContentTypeJsonValue)
	}
}

type service struct {
	api *APIClient
}

type IssuesService service

type PullRequestsService service

type RepositoryService service

type UserService service
