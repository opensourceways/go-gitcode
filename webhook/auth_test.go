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
package webhook

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	payloadData = "{\n  \"note\": \"/ibforuorg/community-test/pulls/2#note_30974945\" \n }"
)

func TestGitCodeAuthenticationAuth(t *testing.T) {

	type args struct {
		r   GitCodeAuthentication
		w   http.ResponseWriter
		req *http.Request
	}

	testCases := []struct {
		no  string
		in  args
		out string
		fn  func(i *args)
	}{{
		"case0",
		args{
			GitCodeAuthentication{},
			httptest.NewRecorder(),
			func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/case0", nil)
				return req
			}(),
		},
		headerUserAgentErrorMessage,
		nil,
	},
		{
			"case1",
			args{
				GitCodeAuthentication{},
				httptest.NewRecorder(),
				func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/case1", nil)
					req.Header.Set(headerUserAgent, headerUserAgentValue)
					return req
				}(),
			},
			http.StatusText(http.StatusMethodNotAllowed),
			nil,
		},
		{
			"case2",
			args{
				GitCodeAuthentication{},
				httptest.NewRecorder(),
				func() *http.Request {
					req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/case2", nil)
					req.Header.Set(headerUserAgent, headerUserAgentValue)
					return req
				}(),
			},
			headerContentTypeErrorMessage,
			nil,
		},
		{
			"case3",
			args{
				GitCodeAuthentication{},
				httptest.NewRecorder(),
				func() *http.Request {
					req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/case3", nil)
					req.Header.Set(headerUserAgent, headerUserAgentValue)
					req.Header.Set(headerContentTypeName, headerContentTypeJsonValue)
					return req
				}(),
			},
			headerEventErrorMessage,
			nil,
		},
		{
			"case4",
			args{
				GitCodeAuthentication{},
				httptest.NewRecorder(),
				func() *http.Request {
					req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/case4", nil)
					req.Header.Set(headerUserAgent, headerUserAgentValue)
					req.Header.Set(headerContentTypeName, headerContentTypeJsonValue)
					req.Header.Set(headerEventType, noteEvent)
					return req
				}(),
			},
			headerEmptyTokenErrorMessage,
			nil,
		},
		{
			"case5",
			args{
				GitCodeAuthentication{signKey: "1234"},
				httptest.NewRecorder(),
				func() *http.Request {
					req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/case5", func() io.Reader {
						var b io.Reader
						b = &bytes.Buffer{}
						return b
					}())
					req.Header.Set(headerUserAgent, headerUserAgentValue)
					req.Header.Set(headerContentTypeName, headerContentTypeJsonValue)
					req.Header.Set(headerEventType, noteEvent)
					req.Header.Set(headerEventToken, "123451")
					return req
				}(),
			},
			headerInvalidTokenErrorMessage,
			nil,
		},
		{
			"case6",
			args{
				GitCodeAuthentication{signKey: "1234"},
				httptest.NewRecorder(),
				func() *http.Request {
					req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/case6", func() io.Reader {
						var b io.Reader
						b = &bytes.Buffer{}
						return b
					}())
					req.Header.Set(headerUserAgent, headerUserAgentValue)
					req.Header.Set(headerContentTypeName, headerContentTypeJsonValue)
					req.Header.Set(headerEventType, noteEvent)
					req.Header.Set(headerEventToken, "sha256=36acf017ea0974457577506ef75268ac93ed6d61864ee994f438b63916ed1736")
					return req
				}(),
			},
			"",
			func(i *args) {
				assert.Equal(t, noteEvent, i.r.eventType)
				assert.Equal(t, *i.r.payload, bytes.Buffer{})
				assert.Equal(t, "1234", i.r.signKey)
			},
		},
		{
			"case7",
			args{
				GitCodeAuthentication{signKey: "1234"},
				httptest.NewRecorder(),
				func() *http.Request {
					req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/case7", func() io.Reader {
						var b io.Reader
						buf := &bytes.Buffer{}
						buf.Write([]byte(payloadData))
						b = buf
						return b
					}())
					req.Header.Set(headerUserAgent, headerUserAgentValue)
					req.Header.Set(headerContentTypeName, headerContentTypeJsonValue)
					req.Header.Set(headerEventType, noteEvent)
					req.Header.Set(headerEventToken, "sha256=f585860d0ca237e0550da0e166370b9c372e8aeb2e639b0ac9884cd52681c576")
					return req
				}(),
			},
			"",
			func(i *args) {
				assert.Equal(t, noteEvent, i.r.eventType)
				if i.r.payload == nil {
					t.Error("payload should be non-nil")
				}
				assert.Equal(t, i.r.payload.String(), payloadData)
				assert.Equal(t, "1234", i.r.signKey)
			},
		},
		{
			"case8",
			args{
				GitCodeAuthentication{signKey: "1234"},
				httptest.NewRecorder(),
				func() *http.Request {
					req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/case8", func() io.Reader {
						var b io.Reader
						buf := &bytes.Buffer{}
						buf.Write([]byte(payloadData))
						b = buf
						return b
					}())
					req.Header.Set(headerUserAgent, headerUserAgentValue)
					req.Header.Set(headerContentTypeName, headerContentTypeJsonValue)
					req.Header.Set(headerEventType, noteEvent)
					req.Header.Set(headerEventToken, "sha256=36acf017ea0974457577506ef75268ac93ed6d61864ee994f438b63916ed1736")

					_, _ = io.Copy(io.Discard, req.Body)
					return req
				}(),
			},
			"",
			func(i *args) {
				assert.Equal(t, "", i.r.payload.String())
			},
		},
		{
			"case9",
			args{
				GitCodeAuthentication{signKey: "1234"},
				httptest.NewRecorder(),
				nil,
			},
			errorNilRequest.Error(),
			nil,
		},
	}

	for i := range testCases {
		t.Run(testCases[i].no, func(t *testing.T) {
			got, _ := testCases[i].in.r.Auth(testCases[i].in.w, testCases[i].in.req)
			if got != nil {
				assert.Equal(t, testCases[i].out, got.Error())
			}
			if testCases[i].fn != nil {
				testCases[i].fn(&testCases[i].in)
			}
		})
	}
}

func TestGitCodeAuthenticationAuthByMock(t *testing.T) {
	e := errors.New("fad")
	patch := gomonkey.ApplyFunc(ReadPayload, func(w http.ResponseWriter, r *http.Request) (*bytes.Buffer, error) {
		return nil, e
	})

	defer patch.Reset()

	a := GitCodeAuthentication{signKey: "1234"}
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/case8", nil)
	req.Header.Set(headerUserAgent, headerUserAgentValue)

	err, _ := a.Auth(httptest.NewRecorder(), req)
	assert.Equal(t, err, e)
}

func TestGitCodeAuthenticationSetSignKey(t *testing.T) {

	type args struct {
		r     GitCodeAuthentication
		token []byte
	}

	testCases := []struct {
		no  string
		in  args
		out error
		fn  func(i *args)
	}{
		{
			"case1",
			args{
				GitCodeAuthentication{},
				nil,
			},
			errorNilToken,
			nil,
		},
		{
			"case2",
			args{
				GitCodeAuthentication{},
				[]byte(""),
			},
			errorNilToken,
			nil,
		},
		{
			"case3",
			args{
				GitCodeAuthentication{},
				[]byte("12345"),
			},
			nil,
			func(i *args) {
				assert.Equal(t, []byte("12345"), i.token)
			},
		},
		{
			"case4",
			args{
				GitCodeAuthentication{},
				[]byte("gfasdihgo;pogfjaklhsbd"),
			},
			nil,
			func(i *args) {
				assert.Equal(t, []byte("gfasdihgo;pogfjaklhsbd"), i.token)
			},
		},
	}

	for i := range testCases {
		t.Run(testCases[i].no, func(t *testing.T) {
			got := testCases[i].in.r.SetSignKey(testCases[i].in.token)
			assert.Equal(t, testCases[i].out, got)
			if testCases[i].fn != nil {
				testCases[i].fn(&testCases[i].in)
			}
		})
	}
}

func TestGitCodeAuthenticationGetPayload(t *testing.T) {

	type args struct {
		r GitCodeAuthentication
	}
	b := &bytes.Buffer{}
	b.Write([]byte("aaadssd"))

	testCases := []struct {
		no  string
		in  args
		out *bytes.Buffer
		fn  func(i *args)
	}{
		{
			"case1",
			args{
				GitCodeAuthentication{},
			},
			nil,
			nil,
		},
		{
			"case2",
			args{
				GitCodeAuthentication{
					payload: b,
				},
			},
			b,
			func(i *args) {
				assert.Equal(t, i.r.payload.String(), "aaadssd")
			},
		},
	}

	for i := range testCases {
		t.Run(testCases[i].no, func(t *testing.T) {
			got := testCases[i].in.r.GetPayload()
			assert.Equal(t, testCases[i].out, got)
			if testCases[i].fn != nil {
				testCases[i].fn(&testCases[i].in)
			}
		})
	}
}

func TestGitCodeAuthenticationGetEventType(t *testing.T) {

	type args struct {
		r GitCodeAuthentication
	}

	testCases := []struct {
		no  string
		in  args
		out string
		fn  func(i *args)
	}{
		{
			"case1",
			args{
				GitCodeAuthentication{},
			},
			"",
			nil,
		},
		{
			"case2",
			args{
				GitCodeAuthentication{
					eventType: "",
				},
			},
			"",
			nil,
		},
		{
			"case3",
			args{
				GitCodeAuthentication{
					eventType: "1234hfas",
				},
			},
			"1234hfas",
			nil,
		},
	}

	for i := range testCases {
		t.Run(testCases[i].no, func(t *testing.T) {
			got := testCases[i].in.r.GetEventType()
			assert.Equal(t, testCases[i].out, got)
			if testCases[i].fn != nil {
				testCases[i].fn(&testCases[i].in)
			}
		})
	}
}

func TestGitCodeAuthenticationGetEventGUID(t *testing.T) {

	type args struct {
		r GitCodeAuthentication
	}

	testCases := []struct {
		no  string
		in  args
		out string
		fn  func(i *args)
	}{
		{
			"case1",
			args{
				GitCodeAuthentication{},
			},
			"",
			nil,
		},
		{
			"case2",
			args{
				GitCodeAuthentication{
					eventGUID: "",
				},
			},
			"",
			nil,
		},
		{
			"case3",
			args{
				GitCodeAuthentication{
					eventGUID: "1234hfas",
				},
			},
			"1234hfas",
			nil,
		},
	}

	for i := range testCases {
		t.Run(testCases[i].no, func(t *testing.T) {
			got := testCases[i].in.r.GetEventGUID()
			assert.Equal(t, testCases[i].out, got)
			if testCases[i].fn != nil {
				testCases[i].fn(&testCases[i].in)
			}
		})
	}
}

func TestGitCodeAuthenticationHandleErr(t *testing.T) {

	assert.Equal(t, fmt.Errorf(httpStatusCodeIncorrectErrorFormat, http.StatusAccepted), handleErr(httptest.NewRecorder(), http.StatusAccepted, ""))
	assert.Equal(t, errorNilResponse, handleErr(nil, http.StatusBadRequest, ""))

	w := httptest.NewRecorder()
	assert.Equal(t, "1234", handleErr(w, http.StatusBadRequest, "1234").Error())
	assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
	var got strings.Builder
	_, _ = io.Copy(&got, w.Result().Body)
	assert.Equal(t, "1234\n", got.String())
}

func TestReadPayload(t *testing.T) {
	e := errors.New("read err")
	patch := gomonkey.ApplyFunc(io.Copy, func(dst io.Writer, src io.Reader) (written int64, err error) {
		return 0, e
	})
	defer patch.Reset()

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/case9", func() io.Reader {
		var b io.Reader
		buf := &bytes.Buffer{}
		buf.Write([]byte(payloadData))
		b = buf
		return b
	}())
	payload, err1 := ReadPayload(httptest.NewRecorder(), req)
	var p *bytes.Buffer
	assert.Equal(t, p, payload)
	assert.Equal(t, e, err1)

}

func TestSignSuccess(t *testing.T) {
	type args struct {
		token   string
		signKey string
		payload *bytes.Buffer
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test with valid token",
			args: args{
				token:   "sha256=3938a65bf0a111e17a7dfe928ea0c73a38c5af80006a939a81c46b372e8f8815",
				signKey: "secret",
				payload: bytes.NewBufferString("{\n    \"content\": \"MTI0MTQxMjQxMjQ=\",\n    \"message\": \"fas\",\n    \"branch\": \"test1-patch-1\"\n}"),
			},
			want: true,
		},
		{
			name: "Test with invalid token prefix",
			args: args{
				token:   "md5=123456",
				signKey: "secret",
				payload: bytes.NewBufferString("test payload"),
			},
			want: false,
		},
		{
			name: "Test with empty token",
			args: args{
				token:   "",
				signKey: "secret",
				payload: bytes.NewBufferString("test payload"),
			},
			want: false,
		},
		{
			name: "Test with empty payload",
			args: args{
				token:   "sha256=123456",
				signKey: "secret",
				payload: bytes.NewBufferString(""),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := signSuccess(tt.args.token, tt.args.signKey, tt.args.payload); got != tt.want {
				t.Errorf("signSuccess() = %v, want %v", got, tt.want)
			}
		})
	}
}
