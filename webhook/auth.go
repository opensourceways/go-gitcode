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
package webhook

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

type GitCodeAuthentication struct {
	payload   *bytes.Buffer
	EventType string
	EventGUID string
	signKey   string
}

func (a *GitCodeAuthentication) GetPayload() *bytes.Buffer {
	return a.payload
}

func (a *GitCodeAuthentication) GetEventType() *string {
	return &a.EventType
}

func (a *GitCodeAuthentication) GetEventGUID() *string {
	return &a.EventGUID
}

func (a *GitCodeAuthentication) Auth(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	var payload bytes.Buffer
	if _, err := io.Copy(&payload, r.Body); err != nil {
		http.Error(w, "400 Bad Request: Failed to read request body", http.StatusBadRequest)
		return err
	}

	a.payload = &payload

	// Header checks: It must be a POST with an event type and a signature.
	if r.Method != http.MethodPost {
		return handleErr(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}

	if v := r.Header.Get("Content-Type"); v != "application/json" {
		return handleErr(w, http.StatusBadRequest, "400 Bad Request: Hook only accepts content-type: application/json")
	}

	if a.EventType = r.Header.Get("X-GitCode-Event"); a.EventType == "" {
		return handleErr(w, http.StatusBadRequest, "400 Bad Request: Missing X-GitCode-Event Header")
	}

	token := r.Header.Get("X-GitCode-Token")
	if token == "" {
		return handleErr(w, http.StatusUnauthorized, "401 Unauthorized: Missing X-GitCode-Token")
	}

	// Validate the payload with our HMAC secret.
	if !signSuccess(token, a.signKey) {
		return handleErr(w, http.StatusUnauthorized, "403 Forbidden: Invalid X-GitCode-Token")
	}

	return nil
}

func handleErr(w http.ResponseWriter, errCode int, errMsg string) error {
	// logging
	http.Error(w, errMsg, errCode)
	return errors.New(errMsg)
}

func signSuccess(token, signKey string) bool {
	return signKey == token
}
