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

func TestAddLabelsToPullRequest(t *testing.T) {
	t.Parallel()
	client, mux, _ := mockServer(t)

	want := new([]*Label)
	_ = readTestdata(t, prTestDataDir+"pull_requests_add_labels.json", want)

	mux.HandleFunc("/repos/"+owner+"/"+repo+"/pulls/33/labels", func(w http.ResponseWriter, r *http.Request) {
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
	t.Parallel()
	client, mux, _ := mockServer(t)

	mux.HandleFunc("/repos/"+owner+"/"+repo+"/pulls/34/labels/fa,fsw", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	ok, err := client.PullRequests.RemoveLabelsFromPullRequest(ctx, owner, repo, "34", []string{"fa", "fsw"})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, ok)
}
