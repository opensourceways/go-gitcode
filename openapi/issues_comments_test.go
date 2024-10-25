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

func TestCreateIssueComment(t *testing.T) {
	t.Parallel()
	client, mux, _ := mockServer(t)

	var comments IssueComment
	_, _ = readTestdata(t, issuesTestDataDir+"issues_comment.json", &comments)

	mux.HandleFunc("/repos/"+owner+"/"+repo+"/issues/1/comments", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(HeaderContentTypeName, HeaderContentTypeJsonValue)
		err := json.NewEncoder(w).Encode(comments)
		if err != nil {
			t.Errorf("Issues.ListLabels mock response data error: %v", err)
		}
	})

	comment := "123987u41"
	result, ok, err := client.Issues.CreateIssueComment(context.Background(), owner, repo, "1", &IssueComment{
		Body: &comment,
	})
	if err != nil {
		t.Errorf("Issues.CreateIssueComment returned error: %v", err)
	}
	assert.Equal(t, true, ok)
	assert.Equal(t, comments, *result)

}
