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

func TestGitCodeAuthenticationAuth(t *testing.T) {
	t.Parallel()

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
	}{
		{
			"case1",
			args{
				GitCodeAuthentication{},
				httptest.NewRecorder(),
				func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/case1", nil)
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
					req.Header.Set("Content-Type", "application/json")
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
					req.Header.Set("Content-Type", "application/json")
					req.Header.Set("X-GitCode-Event", "Note Hook")
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
					req.Header.Set("Content-Type", "application/json")
					req.Header.Set("X-GitCode-Event", "Note Hook")
					req.Header.Set("X-GitCode-Token", "12345")
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
					req.Header.Set("Content-Type", "application/json")
					req.Header.Set("X-GitCode-Event", "Note Hook")
					req.Header.Set("X-GitCode-Token", "1234")
					return req
				}(),
			},
			"",
			func(i *args) {
				assert.Equal(t, "Note Hook", i.r.eventType)
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
						buf.Write([]byte("{\n  \"note\": \"/ibforuorg/community-test/pulls/2#note_30974945\" \n }"))
						b = buf
						return b
					}())
					req.Header.Set("Content-Type", "application/json")
					req.Header.Set("X-GitCode-Event", "Note Hook")
					req.Header.Set("X-GitCode-Token", "1234")
					return req
				}(),
			},
			"",
			func(i *args) {
				assert.Equal(t, "Note Hook", i.r.eventType)
				if i.r.payload == nil {
					t.Error("payload should be non-nil")
				}
				assert.Equal(t, i.r.payload.String(), "{\n  \"note\": \"/ibforuorg/community-test/pulls/2#note_30974945\" \n }")
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
						buf.Write([]byte("{\n  \"note\": \"/ibforuorg/community-test/pulls/2#note_30974945\" \n }"))
						b = buf
						return b
					}())
					req.Header.Set("Content-Type", "application/json")
					req.Header.Set("X-GitCode-Event", "Note Hook")
					req.Header.Set("X-GitCode-Token", "1234")

					_, _ = io.Copy(io.Discard, req.Body)
					return req
				}(),
			},
			"",
			func(i *args) {
				assert.Equal(t, i.r.payload.String(), "")
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
			got := testCases[i].in.r.Auth(testCases[i].in.w, testCases[i].in.req)
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

	err := a.Auth(httptest.NewRecorder(), req)
	assert.Equal(t, err, e)
}

func TestGitCodeAuthenticationSetSignKey(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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

func TestGitCodeAuthenticationsignSuccess(t *testing.T) {
	t.Parallel()

	assert.Equal(t, false, signSuccess("", " "))
	assert.Equal(t, true, signSuccess("", ""))
	assert.Equal(t, true, signSuccess("1231", "1231"))
}

func TestGitCodeAuthenticationhandleErr(t *testing.T) {
	t.Parallel()

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
		buf.Write([]byte("{\n  \"note\": \"/ibforuorg/community-test/pulls/2#note_30974945\" \n }"))
		b = buf
		return b
	}())
	payload, err1 := ReadPayload(httptest.NewRecorder(), req)
	var p *bytes.Buffer
	assert.Equal(t, p, payload)
	assert.Equal(t, e, err1)

}
