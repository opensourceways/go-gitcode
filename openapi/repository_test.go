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
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetRepoContributors(t *testing.T) {

	client, mux, _ := mockServer(t)

	want := new([]*Contributor)
	_ = readTestdata(t, reposTestDataDir+"repository_contributors.json", want)

	mux.HandleFunc(prefixUrlPath+owner+"/"+repo+"/contributors", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		_ = json.NewEncoder(w).Encode(want)

		if r.URL.RawQuery != "" {
			assert.Equal(t, r.URL.RawQuery, "type=jghjhdas")
		}
	})

	ctx := context.Background()
	got, ok, err := client.Repository.GetRepoContributors(ctx, owner, repo, "")
	assert.Equal(t, nil, err)
	assert.Equal(t, true, ok)

	for i := range *want {
		d1, _ := json.Marshal(*(*want)[i])
		d2, _ := json.Marshal(*got[i])
		assert.Equal(t, d1, d2)
	}

	got, ok, err = client.Repository.GetRepoContributors(ctx, owner, repo, "jghjhdas")
	assert.Equal(t, nil, err)
	assert.Equal(t, true, ok)

	for i := range *want {
		d1, _ := json.Marshal(*(*want)[i])
		d2, _ := json.Marshal(*got[i])
		assert.Equal(t, d1, d2)
	}
}
