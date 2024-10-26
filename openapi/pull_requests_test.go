// Copyright 2024 Chao Feng
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

func TestGetPullRequest(t *testing.T) {
	t.Parallel()
	client, mux, _ := mockServer(t)

	want := new(PullRequest)
	_ = readTestdata(t, prTestDataDir+"pull_requests.json", want)

	mux.HandleFunc("/repos/"+owner+"/"+repo+"/pulls/11", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		_ = json.NewEncoder(w).Encode(want)
	})

	ctx := context.Background()
	got, ok, err := client.PullRequests.GetPullRequest(ctx, owner, repo, "11")
	assert.Equal(t, nil, err)
	assert.Equal(t, true, ok)
	d1, _ := json.Marshal(*want)
	d2, _ := json.Marshal(*got)
	assert.Equal(t, d1, d2)
}

func TestUpdatePullRequest(t *testing.T) {
	t.Parallel()
	client, mux, _ := mockServer(t)

	want := new(PullRequest)
	_ = readTestdata(t, prTestDataDir+"pull_requests_closed.json", want)

	mux.HandleFunc("/repos/"+owner+"/"+repo+"/pulls/12", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		_ = json.NewEncoder(w).Encode(want)
	})

	ctx := context.Background()
	got, ok, err := client.PullRequests.UpdatePullRequest(ctx, owner, repo, "12", &PullRequestRequest{
		State: "closed",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, ok)
	d1, _ := json.Marshal(*want)
	d2, _ := json.Marshal(*got)
	assert.Equal(t, d1, d2)

	want1 := new(PullRequest)
	_ = readTestdata(t, prTestDataDir+"pull_requests_open.json", want1)

	mux.HandleFunc("/repos/"+owner+"/"+repo+"/pulls/13", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		_ = json.NewEncoder(w).Encode(want1)
	})

	got1, ok, err := client.PullRequests.UpdatePullRequest(context.Background(), owner, repo, "13", &PullRequestRequest{
		State: "open",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, ok)
	d3, _ := json.Marshal(*want1)
	d4, _ := json.Marshal(*got1)
	assert.Equal(t, d3, d4)
}

func TestListPullRequestLinkingIssues(t *testing.T) {
	t.Parallel()
	client, mux, _ := mockServer(t)

	want := new([]*Issue)
	_ = readTestdata(t, prTestDataDir+"pull_requests_linking_issues.json", want)

	mux.HandleFunc("/repos/"+owner+"/"+repo+"/pulls/15/issues", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		_ = json.NewEncoder(w).Encode(want)
	})

	ctx := context.Background()
	got, ok, err := client.PullRequests.ListPullRequestLinkingIssues(ctx, owner, repo, "15")
	assert.Equal(t, nil, err)
	assert.Equal(t, true, ok)

	for i := range *want {
		d1, _ := json.Marshal(*(*want)[i])
		d2, _ := json.Marshal(*got[i])
		assert.Equal(t, d1, d2)
	}
}
