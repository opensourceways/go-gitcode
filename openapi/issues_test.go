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

func TestUpdateIssue(t *testing.T) {
	t.Parallel()
	client, mux, _ := mockServer(t)

	issue := new(Issue)
	_ = readTestdata(t, issuesTestDataDir+"issues_update.json", issue)

	mux.HandleFunc("/repos/"+owner+"/issues/1", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		err := json.NewEncoder(w).Encode(issue)
		if err != nil {
			t.Errorf("Issues.UpdateIssue mock response data error: %v", err)
		}
	})

	ctx := context.Background()
	result, ok, err := client.Issues.UpdateIssue(ctx, owner, "1", &IssueRequest{
		Repository: repo,
		Title:      "issue1",
	})
	if err != nil {
		t.Errorf("Issues.UpdateIssue returned error: %v", err)
	}
	assert.Equal(t, true, ok)
	assert.Equal(t, issue, result)

	errMsg := "{\n    \"error_code\": 403,\n    \"error_code_name\": \"FORBIDDEN\",\n    \"error_message\": \"no scopes:read_projects\",\n    \"trace_id\": \"33809e888a654b78bb2be8e7c97c9423\"\n}"
	mux.HandleFunc("/repos/"+owner+"/issues/2", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, errMsg, http.StatusBadRequest)
	})

	result, ok, err = client.Issues.UpdateIssue(context.Background(), owner, "2", &IssueRequest{
		Repository: repo,
		Title:      "issue2",
	})

	assert.Equal(t, false, ok)
	assert.Equal(t, Issue{}, *result)
	assert.Equal(t, errMsg+"\n", err.Error())
}

func TestListIssueLinkingPullRequests(t *testing.T) {
	t.Parallel()
	client, mux, _ := mockServer(t)

	prs := new([]*PullRequest)
	_ = readTestdata(t, issuesTestDataDir+"issues_linking_prs.json", prs)

	mux.HandleFunc("/repos/"+owner+"/issues/1/pull_requests", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		err := json.NewEncoder(w).Encode(prs)
		if err != nil {
			t.Errorf("Issues.ListIssueLinkingPullRequests mock response data error: %v", err)
		}
	})

	ctx := context.Background()
	result, ok, err := client.Issues.ListIssueLinkingPullRequests(ctx, owner, repo, "1")
	if err != nil {
		t.Errorf("Issues.ListIssueLinkingPullRequests returned error: %v", err)
	}
	assert.Equal(t, true, ok)
	for i := range *prs {
		d1, _ := json.Marshal(*(*prs)[i])
		d2, _ := json.Marshal(*result[i])
		assert.Equal(t, d1, d2)
	}
}
