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
package openapi_test

import (
	"context"
	"encoding/json"
	"github.com/opensourceways/go-gitcode/gitcode/openapi"
	"net/http"
	"os"
	"testing"
)

func TestListLabels(t *testing.T) {
	t.Parallel()
	client, mux, _ := mockServer(t)

	owner, repo := "111", "222"

	data, _ := os.ReadFile("./testdata/issues/list-labels.json")
	var srcLabels []*openapi.Label
	err := json.Unmarshal(data, &srcLabels)
	if err != nil {
		t.Errorf("Issues.ListLabels mock response data error: %v", err)
	}

	mux.HandleFunc("/repos/"+owner+"/"+repo+"/labels", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, r, http.MethodGet)
		assertReqBody(t, r, http.NoBody)
		w.Header().Set(openapi.HeaderContentTypeName, openapi.HeaderContentTypeJsonValue)
		err = json.NewEncoder(w).Encode(srcLabels)
		if err != nil {
			t.Errorf("Issues.ListLabels mock response data error: %v", err)
		}
	})

	ctx := context.Background()
	targetLabels, _, err := client.Issues.ListRepoIssueLabels(ctx, owner, repo)
	if err != nil {
		t.Errorf("Issues.ListLabels returned error: %v", err)
	}

	assertDataLa(t, srcLabels, targetLabels)
}
