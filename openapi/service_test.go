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
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const (
	owner = "111"
	repo  = "222"

	testDataDir       = "testdata"
	issuesTestDataDir = testDataDir + string(os.PathSeparator) + "issues" + string(os.PathSeparator)
	prTestDataDir     = testDataDir + string(os.PathSeparator) + "pr" + string(os.PathSeparator)
	reposTestDataDir  = testDataDir + string(os.PathSeparator) + "repos" + string(os.PathSeparator)
)

// setup sets up a test HTTP server along with a github.api that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the api method being tested.
func mockServer(t *testing.T) (client *APIClient, mux *http.ServeMux, serverURL string) {
	t.Helper()
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	handlerPath := "/api/v5/"
	apiHandler.Handle(handlerPath, http.StripPrefix(handlerPath[:len(handlerPath)-1], mux))

	// server is a test HTTP server used to provide mock api responses.
	server := httptest.NewServer(apiHandler)

	// api is the GitHub api being tested and is
	// configured to use test server.
	client = NewAPIClientWithAuthorization([]byte("1111111111"))
	uri, _ := url.Parse(server.URL + handlerPath)
	client.BaseURL = uri

	t.Cleanup(server.Close)

	return client, mux, server.URL
}

func readTestdata(t *testing.T, path string, ptr any) ([]byte, error) {

retry:
	i := 0
	absPath, err := filepath.Abs(path)
	if err != nil {
		t.Error(absPath + "not found")
		return nil, err
	}
	if _, err = os.Stat(absPath); !os.IsNotExist(err) {
		data, err := os.ReadFile(absPath)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, ptr)
		if err != nil {
			_, _, line, _ := runtime.Caller(1)
			t.Errorf("code line: %d, error: %v", line, err)
		}
		return data, err
	} else {
		i++
		path = ".." + string(os.PathSeparator) + path
		if i <= 3 {
			goto retry
		}
	}

	return nil, err
}
