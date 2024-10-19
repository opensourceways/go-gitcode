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
	"github.com/opensourceways/go-gitcode/gitcode/openapi"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// setup sets up a test HTTP server along with a github.api that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func mockServer(t *testing.T) (client *openapi.APIClient, mux *http.ServeMux, serverURL string) {
	t.Helper()
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// We want to ensure that tests catch mistakes where the endpoint URL is
	// specified as absolute rather than relative. It only makes a difference
	// when there's a non-empty base URL path. So, use that. See issue #752.
	apiHandler := http.NewServeMux()
	handlerPath := "/api/v5/"
	apiHandler.Handle(handlerPath, http.StripPrefix(handlerPath[:len(handlerPath)-1], mux))

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	// api is the GitHub api being tested and is
	// configured to use test server.
	client = openapi.NewAPIClientWithAuthorization([]byte("1111111111"))
	uri, _ := url.Parse(server.URL + handlerPath)
	client.BaseURL = uri

	t.Cleanup(server.Close)

	return client, mux, server.URL
}
