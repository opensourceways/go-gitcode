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
	_ = readTestdata(t, issuesTestDataDir+"issues_list_labels.json", &srcLabels)

	mux.HandleFunc("/repos/"+owner+"/"+repo+"/labels", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
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
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		// a word wrap \n write to response
		http.Error(w, msg, http.StatusInternalServerError)
	})

	targetLabels, ok, err = client.Issues.ListRepoIssueLabels(context.Background(), owner, repo)
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, len(targetLabels))
	assert.Equal(t, msg+"\n", err.Error())

}

func TestCreateRepoIssueLabel(t *testing.T) {
	t.Parallel()
	client, mux, _ := mockServer(t)

	var srcLabels Label
	_ = readTestdata(t, issuesTestDataDir+"issues_create_label.json", &srcLabels)

	mux.HandleFunc("/repos/"+owner+"/"+repo+"/labels", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		err := json.NewEncoder(w).Encode(srcLabels)
		if err != nil {
			t.Errorf("Issues.ListLabels mock response data error: %v", err)
		}
	})

	targetLabels, ok, err := client.Issues.CreateRepoIssueLabel(context.Background(), owner, repo, &Label{Name: "gfa", Color: "#000"})
	if err != nil {
		t.Errorf("Issues.CreateRepoIssueLabel returned error: %v", err)
	}
	assert.Equal(t, true, ok)
	assert.Equal(t, srcLabels, *targetLabels)

}

func TestUpdateRepoIssueLabel(t *testing.T) {
	t.Parallel()
	client, mux, _ := mockServer(t)

	var srcLabels Label
	_ = readTestdata(t, issuesTestDataDir+"issues_create_label.json", &srcLabels)

	mux.HandleFunc("/repos/"+owner+"/"+repo+"/labels/123", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		err := json.NewEncoder(w).Encode(srcLabels)
		if err != nil {
			t.Errorf("Issues.ListLabels mock response data error: %v", err)
		}
	})

	targetLabels, ok, err := client.Issues.UpdateRepoIssueLabel(context.Background(), owner, repo, "123", "gf3", "#001")
	if err != nil {
		t.Errorf("Issues.UpdateRepoIssueLabel returned error: %v", err)
	}
	assert.Equal(t, true, ok)
	assert.Equal(t, srcLabels, *targetLabels)

}

func TestDeleteRepoIssueLabel(t *testing.T) {
	t.Parallel()
	client, mux, _ := mockServer(t)

	mux.HandleFunc("/repos/"+owner+"/"+repo+"/labels/gga3g", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
	})

	ok, err := client.Issues.DeleteRepoIssueLabel(context.Background(), owner, repo, "gga3g")
	if err != nil {
		t.Errorf("Issues.UpdateRepoIssueLabel returned error: %v", err)
	}
	assert.Equal(t, true, ok)

	//var str strings.Builder
	data := readTestdata(t, issuesTestDataDir+"issues_delete_label_failed.json", nil)

	mux.HandleFunc("/repos/"+owner+"/"+repo+"/labels/ccs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write(data)
	})

	ok, err = client.Issues.DeleteRepoIssueLabel(context.Background(), owner, repo, "ccs")
	assert.Equal(t, false, ok)
	assert.Equal(t, string(data), err.Error())

}

func TestAddLabelsToIssue(t *testing.T) {
	t.Parallel()
	client, mux, _ := mockServer(t)

	var want []*Label
	_ = readTestdata(t, issuesTestDataDir+"issues_add_labels.json", &want)

	mux.HandleFunc("/repos/"+owner+"/"+repo+"/issues/fasd/labels", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		w.WriteHeader(http.StatusCreated)
		err := json.NewEncoder(w).Encode(want)
		if err != nil {
			t.Errorf("Issues.ListLabels mock response data error: %v", err)
		}
	})

	got, ok, err := client.Issues.AddLabelsToIssue(context.Background(), owner, repo, "fasd", []string{"858"})
	if err != nil {
		t.Errorf("Issues.UpdateRepoIssueLabel returned error: %v", err)
	}
	assert.Equal(t, true, ok)
	for i := range want {
		assert.Equal(t, *want[i], *got[i])
	}

}

func TestRemoveLabelsFromIssue(t *testing.T) {
	t.Parallel()
	client, mux, _ := mockServer(t)

	mux.HandleFunc("/repos/"+owner+"/"+repo+"/issues/623/labels/0gjds", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		w.WriteHeader(http.StatusNoContent)
	})

	ok, err := client.Issues.RemoveLabelsFromIssue(context.Background(), owner, repo, "623", "0gjds")
	if err != nil {
		t.Errorf("Issues.RemoveLabelsFromIssue returned error: %v", err)
	}
	assert.Equal(t, true, ok)

}
