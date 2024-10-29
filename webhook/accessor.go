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
	"encoding/json"
	"net/http"
)

type GitCodeAccessor struct {
	Issues *IssueEvent
	PR     *PullRequestEvent
	Note   *NoteEvent
}

func (a *GitCodeAccessor) GetAccessor(w http.ResponseWriter, r *http.Request) (any, *bytes.Buffer, *string, *string, bool) {
	payload, err := ReadPayload(w, r)
	if err != nil {
		return nil, nil, nil, nil, false
	}

	eventGUID := r.Header.Get(headerEventGUID)
	eventType := r.Header.Get(headerEventType)

	switch eventType {
	case "Issue Hook":
		a.Issues = new(IssueEvent)
		_ = json.Unmarshal(payload.Bytes(), a.Issues)
		return a.Issues, payload, &eventType, &eventGUID, false
	case "Merge Request Hook":
		a.PR = new(PullRequestEvent)
		_ = json.Unmarshal(payload.Bytes(), a.PR)
		return a.PR, payload, &eventType, &eventGUID, false
	case "Note Hook":
		a.Note = new(NoteEvent)
		_ = json.Unmarshal(payload.Bytes(), a.Note)
		return a.Note, payload, &eventType, &eventGUID, a.Note != nil && a.Note.PR != nil
	default:

	}

	return nil, payload, &eventType, &eventGUID, false
}
