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

func TestListLabels(t *testing.T) {
	t.Parallel()
	client, mux, _ := mockServer(t)

	owner, repo := "111", "222"

	var srcLabels []*Label

	LoadJsonFile(t, "testdata/issues/list-labels.json", &srcLabels)

	mux.HandleFunc("/repos/"+owner+"/"+repo+"/labels", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodGet)
		assert.Equal(t, r.Body, http.NoBody)
		w.Header().Set(HeaderContentTypeName, HeaderContentTypeJsonValue)
		err := json.NewEncoder(w).Encode(srcLabels)
		if err != nil {
			t.Errorf("Issues.ListLabels mock response data error: %v", err)
		}
	})

	ctx := context.Background()
	targetLabels, _, err := client.Issues.ListRepoIssueLabels(ctx, owner, repo)
	if err != nil {
		t.Errorf("Issues.ListLabels returned error: %v", err)
	}

	assert.Equal(t, srcLabels, targetLabels)

}
