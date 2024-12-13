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
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

const (
	webhookTestDataDir = "testdata" + string(os.PathSeparator) + "webhook" + string(os.PathSeparator)
	htmlUrl            = "https://gitcode.com/ibforuorg/test1/issues/4"
)

func TestGetAccessor(t *testing.T) {

	createIssue(t)
	createPR(t)
	notePR(t)
	noteIssue(t)
	pushCode(t)

	buf := &bytes.Buffer{}
	buf.Write([]byte("kjhygadsskhj"))
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/1", buf)
	req.Header.Set(headerEventType, "Dummy Hook")
	req.Header.Set(headerEventGUID, "fasgasd")
	w := httptest.NewRecorder()

	a := new(GitCodeAccessor)
	got1, _, got2, got3 := a.GetAccessor(w, req)
	assert.Equal(t, nil, got1)
	assert.Equal(t, "Dummy Hook", *got2)
	assert.Equal(t, "fasgasd", *got3)
}

func createIssue(t *testing.T) {
	want := GitCodeAccessor{Issues: new(IssueEvent)}
	data := readWebHookTestdata(t, webhookTestDataDir+"issues_create.json", want.Issues)

	buf := &bytes.Buffer{}
	buf.Write(data)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/0", buf)
	req.Header.Set(headerEventType, issueEvent)
	req.Header.Set(headerEventGUID, "1231321")
	w := httptest.NewRecorder()

	a := new(GitCodeAccessor)
	got1, _, got2, got3 := a.GetAccessor(w, req)
	d1, _ := json.Marshal(want.Issues)
	d2, _ := json.Marshal(got1)
	assert.Equal(t, d1, d2)

	assert.Equal(t, issueEvent, *got2)
	assert.Equal(t, "1231321", *got3)

	issue, _ := got1.(*IssueEvent)
	assert.Equal(t, "open", *issue.GetAction())
	assert.Equal(t, (*string)(nil), issue.GetActionDetail())
	assert.Equal(t, "opened", *issue.GetState())
	assert.Equal(t, "ibforuorg", *issue.GetOrg())
	assert.Equal(t, "test1", *issue.GetRepo())
	assert.Equal(t, htmlUrl, *issue.GetHtmlURL())
	assert.Equal(t, "4", *issue.GetNumber())
	assert.Equal(t, "515443", *issue.GetID())
	assert.Equal(t, "*****", *issue.GetAuthor())
	assert.Equal(t, "2024-10-26T10:28:03+08:00", *issue.GetCreateTime())
	assert.Equal(t, "2024-10-26T10:28:03+08:00", *issue.GetUpdateTime())

	issue = new(IssueEvent)
	rt := reflect.TypeOf(issue)
	n := rt.NumMethod()
	for i := 0; i < n; i++ {
		rm := rt.Method(i)
		ret := rm.Func.Call([]reflect.Value{reflect.ValueOf(issue)})
		assert.Equal(t, (*string)(nil), ret[0].Interface())
	}
}

func pushCode(t *testing.T) {
	want := GitCodeAccessor{Push: new(PushEvent)}
	data := readWebHookTestdata(t, webhookTestDataDir+"push_code.json", want.Push)

	buf := &bytes.Buffer{}
	buf.Write(data)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/1", buf)
	req.Header.Set(headerEventType, "Push Hook")
	req.Header.Set(headerEventGUID, "fasgasd")
	w := httptest.NewRecorder()

	a := new(GitCodeAccessor)
	got1, _, got2, got3 := a.GetAccessor(w, req)
	d1, _ := json.Marshal(want.Push)
	d2, _ := json.Marshal(got1)
	assert.Equal(t, d1, d2)

	assert.Equal(t, "Push Hook", *got2)
	assert.Equal(t, "fasgasd", *got3)

	pr, _ := got1.(*PushEvent)
	assert.Equal(t, (*string)(nil), pr.GetAction())
	assert.Equal(t, (*string)(nil), pr.GetState())
	assert.Equal(t, "ibforuorg", *pr.GetOrg())
	assert.Equal(t, "org-repo-role-member-manage", *pr.GetRepo())
	assert.Equal(t, "https://gitcode.com/ibforuorg/org-repo-role-member-manage", *pr.GetHtmlURL())
	assert.Equal(t, "dev", *pr.GetBase())
	assert.Equal(t, "ibforu", *pr.GetAuthor())

	pr = new(PushEvent)
	rt := reflect.TypeOf(pr)
	n := rt.NumMethod()
	for i := 0; i < n; i++ {
		rm := rt.Method(i)
		ret := rm.Func.Call([]reflect.Value{reflect.ValueOf(pr)})
		assert.Equal(t, (*string)(nil), ret[0].Interface())
	}
}

func createPR(t *testing.T) {
	want := GitCodeAccessor{PR: new(PullRequestEvent)}
	data := readWebHookTestdata(t, webhookTestDataDir+"pr_create.json", want.PR)

	buf := &bytes.Buffer{}
	buf.Write(data)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/1", buf)
	req.Header.Set(headerEventType, pullRequestEvent)
	req.Header.Set(headerEventGUID, "fasgasd")
	w := httptest.NewRecorder()

	a := new(GitCodeAccessor)
	got1, _, got2, got3 := a.GetAccessor(w, req)
	d1, _ := json.Marshal(want.PR)
	d2, _ := json.Marshal(got1)
	assert.Equal(t, d1, d2)

	assert.Equal(t, pullRequestEvent, *got2)
	assert.Equal(t, "fasgasd", *got3)

	pr, _ := got1.(*PullRequestEvent)
	assert.Equal(t, "open", *pr.GetAction())
	assert.Equal(t, "", *pr.GetActionDetail())
	assert.Equal(t, "opened", *pr.GetState())
	assert.Equal(t, "ibforuorg", *pr.GetOrg())
	assert.Equal(t, "test1", *pr.GetRepo())
	assert.Equal(t, "https://gitcode.com/ibforuorg/test1/merge_requests/4", *pr.GetHtmlURL())
	assert.Equal(t, "main", *pr.GetBase())
	assert.Equal(t, "ibforuorg/test1/24124124124", *pr.GetHead())
	assert.Equal(t, "4", *pr.GetNumber())
	assert.Equal(t, "190370", *pr.GetID())
	assert.Equal(t, "****", *pr.GetAuthor())
	assert.Equal(t, "2024-10-26T10:32:40+08:00", *pr.GetCreateTime())
	assert.Equal(t, "2024-10-26T10:32:41+08:00", *pr.GetUpdateTime())

	pr = new(PullRequestEvent)
	rt := reflect.TypeOf(pr)
	n := rt.NumMethod()
	for i := 0; i < n; i++ {
		rm := rt.Method(i)
		ret := rm.Func.Call([]reflect.Value{reflect.ValueOf(pr)})
		assert.Equal(t, (*string)(nil), ret[0].Interface())
	}
}

func notePR(t *testing.T) {
	want := GitCodeAccessor{Note: new(NoteEvent)}
	data := readWebHookTestdata(t, webhookTestDataDir+"pr_note.json", want.Note)

	buf := &bytes.Buffer{}
	buf.Write(data)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/2", buf)
	req.Header.Set(headerEventType, noteEvent)
	req.Header.Set(headerEventGUID, "651234123")
	w := httptest.NewRecorder()

	a := new(GitCodeAccessor)
	got1, _, got2, got3 := a.GetAccessor(w, req)
	d1, _ := json.Marshal(want.Note)
	d2, _ := json.Marshal(got1)
	assert.Equal(t, d1, d2)

	assert.Equal(t, noteEvent, *got2)
	assert.Equal(t, "651234123", *got3)

	note, _ := got1.(*NoteEvent)
	assert.Equal(t, "open", *note.GetAction())
	assert.Equal(t, (*string)(nil), note.GetActionDetail())
	assert.Equal(t, "opened", *note.GetState())
	assert.Equal(t, "ibforuorg", *note.GetOrg())
	assert.Equal(t, "test1", *note.GetRepo())
	assert.Equal(t, "https://gitcode.com/ibforuorg/test1/merge_requests/4#note_71e9657489bcddbed4c0a9d2b1e29eb7c8ab26c3", *note.GetHtmlURL())
	assert.Equal(t, "main", *note.GetBase())
	assert.Equal(t, "ibforuorg/test1/24124124124", *note.GetHead())
	assert.Equal(t, "4", *note.GetNumber())
	assert.Equal(t, "190370", *note.GetID())
	assert.Equal(t, "****", *note.GetAuthor())
	assert.Equal(t, "71e9657489bcddbed4c0a9d2b1e29eb7c8ab26c3", *note.GetCommentID())
	assert.Equal(t, "MergeRequest", *note.GetCommentKind())
	assert.Equal(t, "/lgtm\n/approve", *note.GetComment())
	assert.Equal(t, "****", *note.GetCommenter())
	assert.Equal(t, "2024-10-26T11:44:15+08:00", *note.GetCreateTime())
	assert.Equal(t, "2024-10-26T11:44:15+08:00", *note.GetUpdateTime())
}

func noteIssue(t *testing.T) {
	want := GitCodeAccessor{Note: new(NoteEvent)}
	data := readWebHookTestdata(t, webhookTestDataDir+"issues_note.json", want.Note)

	buf := &bytes.Buffer{}
	buf.Write(data)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/2", buf)
	req.Header.Set(headerEventType, noteEvent)
	req.Header.Set(headerEventGUID, "151231321")
	w := httptest.NewRecorder()

	a := new(GitCodeAccessor)
	got1, _, got2, got3 := a.GetAccessor(w, req)
	d1, _ := json.Marshal(want.Note)
	d2, _ := json.Marshal(got1)
	assert.Equal(t, d1, d2)

	assert.Equal(t, noteEvent, *got2)
	assert.Equal(t, "151231321", *got3)

	note, _ := got1.(*NoteEvent)
	assert.Equal(t, (*string)(nil), note.GetAction())
	assert.Equal(t, (*string)(nil), note.GetActionDetail())
	assert.Equal(t, "opened", *note.GetState())
	assert.Equal(t, "ibforuorg", *note.GetOrg())
	assert.Equal(t, "test1", *note.GetRepo())
	assert.Equal(t, "https://gitcode.com/ibforuorg/test1/issues/4#note_d3ab73b290d6fcd8800177e2d34545c755af3af1", *note.GetHtmlURL())
	assert.Equal(t, (*string)(nil), note.GetBase())
	assert.Equal(t, (*string)(nil), note.GetHead())
	assert.Equal(t, "4", *note.GetNumber())
	assert.Equal(t, "515443", *note.GetID())
	assert.Equal(t, "****", *note.GetAuthor())
	assert.Equal(t, "d3ab73b290d6fcd8800177e2d34545c755af3af1", *note.GetCommentID())
	assert.Equal(t, "Issue", *note.GetCommentKind())
	assert.Equal(t, "oiugbfaijub", *note.GetComment())
	assert.Equal(t, "****", *note.GetCommenter())
	assert.Equal(t, "2024-10-26T11:42:05+08:00", *note.GetCreateTime())
	assert.Equal(t, "2024-10-26T11:42:05+08:00", *note.GetUpdateTime())

	note = new(NoteEvent)
	rt := reflect.TypeOf(note)
	n := rt.NumMethod()
	for i := 0; i < n; i++ {
		rm := rt.Method(i)
		ret := rm.Func.Call([]reflect.Value{reflect.ValueOf(note)})
		assert.Equal(t, (*string)(nil), ret[0].Interface())
	}
}

func readWebHookTestdata(t *testing.T, path string, ptr any) []byte {

	i := 0
retry:
	absPath, err := filepath.Abs(path)
	if err != nil {
		t.Error(path + " not found")
		return nil
	}
	if _, err = os.Stat(absPath); !os.IsNotExist(err) {
		data, err := os.ReadFile(absPath)
		if err != nil {
			t.Error(path + " read failed")
			return nil
		}
		if ptr != nil {
			err = json.Unmarshal(data, ptr)
			if err != nil {
				_, _, line, _ := runtime.Caller(1)
				t.Errorf("code line: %d, error: %v", line, err)
			}
		}
		return data
	} else {
		i++
		path = ".." + string(os.PathSeparator) + path
		if i <= 3 {
			goto retry
		}
	}

	t.Error(path + " not found")
	return nil
}
