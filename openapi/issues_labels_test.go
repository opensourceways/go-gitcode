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

	var srcLabels []*Label
	_, _ = readTestdata(t, issuesTestDataDir+"issues_list_labels.json", &srcLabels)

	mux.HandleFunc("/repos/"+owner+"/"+repo+"/labels", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(HeaderContentTypeName, HeaderContentTypeJsonValue)
		err := json.NewEncoder(w).Encode(srcLabels)
		if err != nil {
			t.Errorf("Issues.ListLabels mock response data error: %v", err)
		}
	})

	targetLabels, ok, err := client.Issues.ListRepoIssueLabels(context.Background(), owner, repo)
	if err != nil {
		t.Errorf("Issues.ListLabels returned error: %v", err)
	}
	assert.Equal(t, true, ok)
	assert.Equal(t, srcLabels, targetLabels)

	repo := "333"
	msg := "{\n    \"error_code\": 500,\n    \"error_code_name\": \"FAIL\",\n    \"error_message\": \"系统错误\",\n    \"trace_id\": \"d0834ebae0074f5cab66ef8b8fc529d5\"\n}"
	mux.HandleFunc("/repos/"+owner+"/"+repo+"/labels", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(HeaderContentTypeName, HeaderContentTypeJsonValue)
		// a word wrap \n write to response
		http.Error(w, msg, http.StatusInternalServerError)
	})

	targetLabels, ok, err = client.Issues.ListRepoIssueLabels(context.Background(), owner, repo)
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, len(targetLabels))
	assert.Equal(t, msg+"\n", err.Error())

}
