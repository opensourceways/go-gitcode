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
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"testing"
)

func LoadJsonFile(t *testing.T, path string, ptr any) {

retry:
	absPath, err := filepath.Abs(path)
	if err != nil {
		t.Errorf("%s not found", path)
	}

	stat, err := os.Stat(absPath)
	if err != nil {
		if path[:2] != ".." {
			path = "../" + path
			goto retry
		}

		t.Errorf("get %s info occur err: %v", absPath, err)
		return
	}
	if stat.Mode() != fs.ModeType {
		data, _ := os.ReadFile(absPath)
		err = json.Unmarshal(data, ptr)
		if err != nil {
			t.Errorf("json data convert to struct, occur error: %v", err)
		}
	}
}

func Test_PreOperate(t *testing.T) {
	t.Parallel()

	type args struct {
		h    RequestHandler
		uri  *url.URL
		body any
	}
	testCases := map[string]struct {
		in  args
		out error
		fn  func(i *args)
	}{
		"case1": {
			args{
				RequestHandler{},
				&url.URL{},
				nil,
			},
			nil,
			nil,
		},
		"case2": {
			args{
				RequestHandler{t: Query},
				&url.URL{},
				nil,
			},
			nil,
			func(i *args) {
				assert.Equal(t, i.uri.RawQuery, "")
			},
		},
		"case3": {
			args{
				RequestHandler{t: Query},
				&url.URL{},
				url.Values{},
			},
			nil,
			func(i *args) {
				assert.Equal(t, i.uri.RawQuery, "")
			},
		},
	}

	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tt.in.h.PreOperate(tt.in.uri, tt.in.body)
			assert.Equal(t, tt.out, got)
			if tt.fn != nil {
				tt.fn(&tt.in)
			}
		})
	}
}
