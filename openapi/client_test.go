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
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestParseResp(t *testing.T) {
	w := httptest.NewRecorder()
	resp := w.Result()
	resp1, err := parseResp(resp, nil)
	assert.Equal(t, resp, resp1)
	assert.Equal(t, nil, err)
	//assert.Equal(t, nilContentError, err)

	r1 := struct{}{}
	resp1, err = parseResp(resp, r1)
	assert.Equal(t, resp, resp1)
	assert.Equal(t, respReceiverNotAnPointerError, err)

	r2 := new(struct{})
	resp1, err = parseResp(resp, r2)
	assert.Equal(t, resp, resp1)
	assert.Equal(t, nil, err)
	assert.Equal(t, struct{}{}, *r2)

	type simple struct {
		A string `json:"a"`
		B int    `json:"b"`
	}
	str := "{\n    \"a\": \"4123\",\n    \"b\": 912\n}"
	w1 := httptest.NewRecorder()
	_, _ = w1.Write([]byte(str))
	resp3 := w1.Result()

	r4 := new(simple)
	resp4, err := parseResp(resp3, r4)
	assert.Equal(t, resp3, resp4)
	assert.Equal(t, nil, err)
	assert.Equal(t, simple{A: "4123", B: 912}, *r4)
}
