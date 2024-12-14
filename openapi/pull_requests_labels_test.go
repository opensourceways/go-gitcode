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

func TestAddLabelsToPullRequest(t *testing.T) {

	client, mux, _ := mockServer(t)

	want := new([]*Label)
	_ = readTestdata(t, prTestDataDir+"pull_requests_add_labels.json", want)

	mux.HandleFunc(prefixUrlPath+owner+"/"+repo+"/pulls/33/labels", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		_ = json.NewEncoder(w).Encode(want)
	})

	ctx := context.Background()
	got, ok, err := client.PullRequests.AddLabelsToPullRequest(ctx, owner, repo, "33", []string{"fa", "fsw"})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, ok)

	d1, _ := json.Marshal(want)
	d2, _ := json.Marshal(got)
	assert.Equal(t, d1, d2)
}

func TestRemoveLabelsFromPullRequest(t *testing.T) {

	client, mux, _ := mockServer(t)

	mux.HandleFunc(prefixUrlPath+owner+"/"+repo+"/pulls/34/labels/fa,fsw", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	ok, err := client.PullRequests.RemoveLabelsFromPullRequest(ctx, owner, repo, "34", []string{"fa", "fsw"})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, ok)
}

func TestGetLabelsOfPullRequest(t *testing.T) {

	client, mux, _ := mockServer(t)

	var labels []*Label
	_ = readTestdata(t, prTestDataDir+"pull_requests_add_labels.json", &labels)

	mux.HandleFunc(prefixUrlPath+owner+"/"+repo+"/pulls/4432/labels", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		err := json.NewEncoder(w).Encode(labels)
		if err != nil {
			t.Errorf("PR.GetLabelsOfPullRequest mock response data error: %v", err)
		}
	})

	result, ok, err := client.PullRequests.GetLabelsOfPullRequest(context.Background(), owner, repo, "4432")
	if err != nil {
		t.Errorf("PR.GetLabelsOfPullRequest returned error: %v", err)
	}
	assert.Equal(t, true, ok)
	for i := range labels {
		assert.Equal(t, *labels[i], *result[i])
	}

}
