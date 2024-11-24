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

func TestGetRepoAllMember(t *testing.T) {

	client, mux, _ := mockServer(t)

	want := new([]*User)
	_ = readTestdata(t, reposTestDataDir+"repository_members.json", want)

	mux.HandleFunc(prefixUrlPath+owner+"/"+repo+"/collaborators", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		_ = json.NewEncoder(w).Encode(want)

		if r.URL.RawQuery != "" {
			assert.Equal(t, r.URL.RawQuery, "page=1&per_page=100")
		}
	})

	ctx := context.Background()
	got, ok, err := client.Repository.GetRepoAllMember(ctx, owner, repo, "1")
	assert.Equal(t, nil, err)
	assert.Equal(t, true, ok)

	for i := range *want {
		d1, _ := json.Marshal(*(*want)[i])
		d2, _ := json.Marshal(*got[i])
		assert.Equal(t, d1, d2)
	}

}

func TestGetRepoMemberPermission(t *testing.T) {

	client, mux, _ := mockServer(t)

	user := "{\n    \"id\": 412,\n    \"login\": \"fasfa\",\n    \"permission\": \"admin\"\n}"
	mux.HandleFunc(prefixUrlPath+owner+"/"+repo+"/collaborators/fasfa/permission", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(user))
	})

	ctx := context.Background()
	got, ok, err := client.Repository.GetRepoMemberPermission(ctx, owner, repo, "fasfa")
	assert.Equal(t, nil, err)
	assert.Equal(t, [2]bool{true, false}, ok)
	assert.Equal(t, "fasfa", *got.Login)
	assert.Equal(t, "admin", *got.Permission)

	msg := "{\"message\":\"404 Not Found\"}"
	mux.HandleFunc(prefixUrlPath+owner+"/"+repo+"/collaborators/145123/permission", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(msg))
	})

	ctx1 := context.Background()
	got, ok, err = client.Repository.GetRepoMemberPermission(ctx1, owner, repo, "145123")
	assert.Equal(t, msg, err.Error())
	assert.Equal(t, [2]bool{false, true}, ok)
	assert.Equal(t, User{}, *got)
}

func TestCheckUserIsRepoMember(t *testing.T) {

	client, mux, _ := mockServer(t)

	mux.HandleFunc(prefixUrlPath+owner+"/"+repo+"/collaborators/fasdagsdf", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	got, ok, err := client.Repository.CheckUserIsRepoMember(ctx, owner, repo, "fasdagsdf")
	assert.Equal(t, nil, err)
	assert.Equal(t, false, ok)
	assert.Equal(t, true, got)

	msg := "{\"message\":\"404 Not Found\"}"
	mux.HandleFunc(prefixUrlPath+owner+"/"+repo+"/collaborators/63453", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeName, headerContentTypeJsonValue)
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(msg))
	})

	ctx1 := context.Background()
	got, ok, err = client.Repository.CheckUserIsRepoMember(ctx1, owner, repo, "63453")
	assert.Equal(t, msg, err.Error())
	assert.Equal(t, true, ok)
	assert.Equal(t, false, got)
}
