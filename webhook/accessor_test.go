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
package webhook

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const (
	webhookTestDataDir = "testdata" + string(os.PathSeparator) + "webhook" + string(os.PathSeparator)
)

func TestGetAccessor(t *testing.T) {
	t.Parallel()

	createIssue(t)
	createPR(t)
	notePR(t)
	noteIssue(t)
}

func createIssue(t *testing.T) {
	want := GitCodeAccessor{Issues: new(IssueEvent)}
	data := readWebHookTestdata(t, webhookTestDataDir+"issues_create.json", want.Issues)

	buf := &bytes.Buffer{}
	buf.Write(data)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/0", buf)
	req.Header.Set(headerEventType, "issue_hooks")
	req.Header.Set(headerEventGUID, "1231321")
	w := httptest.NewRecorder()

	a := new(GitCodeAccessor)
	got1, got2, got3, got4 := a.GetAccessor(w, req)
	d1, _ := json.Marshal(want.Issues)
	d2, _ := json.Marshal(got1)
	assert.Equal(t, d1, d2)

	assert.Equal(t, "issue_hooks", *got2)
	assert.Equal(t, "1231321", *got3)
	assert.Equal(t, false, got4)

	issue, _ := got1.(*IssueEvent)
	assert.Equal(t, "open", *issue.GetAction())
	assert.Equal(t, "opened", *issue.GetState())
	assert.Equal(t, "ibforuorg", *issue.GetOrg())
	assert.Equal(t, "test1", *issue.GetRepo())
	assert.Equal(t, "https://gitcode.com/ibforuorg/test1", *issue.GetHtmlURL())
	assert.Equal(t, (*string)(nil), issue.GetBase())
	assert.Equal(t, (*string)(nil), issue.GetHead())
	assert.Equal(t, "4", *issue.GetNumber())
	assert.Equal(t, "ibforu", *issue.GetAuthor())
	assert.Equal(t, (*string)(nil), issue.GetComment())
	assert.Equal(t, (*string)(nil), issue.GetCommenter())
	assert.Equal(t, 0, len(issue.ListLabels()))

	issue = new(IssueEvent)
	assert.Equal(t, (*string)(nil), issue.GetAction())
	assert.Equal(t, (*string)(nil), issue.GetState())
	assert.Equal(t, (*string)(nil), issue.GetOrg())
	assert.Equal(t, (*string)(nil), issue.GetRepo())
	assert.Equal(t, (*string)(nil), issue.GetHtmlURL())
	assert.Equal(t, (*string)(nil), issue.GetBase())
	assert.Equal(t, (*string)(nil), issue.GetHead())
	assert.Equal(t, (*string)(nil), issue.GetNumber())
	assert.Equal(t, (*string)(nil), issue.GetAuthor())
	assert.Equal(t, (*string)(nil), issue.GetComment())
	assert.Equal(t, (*string)(nil), issue.GetCommenter())
	assert.Equal(t, 0, len(issue.ListLabels()))
}

func createPR(t *testing.T) {
	want := GitCodeAccessor{PR: new(PullRequestEvent)}
	data := readWebHookTestdata(t, webhookTestDataDir+"pr_create.json", want.PR)

	buf := &bytes.Buffer{}
	buf.Write(data)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/1", buf)
	req.Header.Set(headerEventType, "merge_request_hooks")
	req.Header.Set(headerEventGUID, "fasgasd")
	w := httptest.NewRecorder()

	a := new(GitCodeAccessor)
	got1, got2, got3, got4 := a.GetAccessor(w, req)
	d1, _ := json.Marshal(want.PR)
	d2, _ := json.Marshal(got1)
	assert.Equal(t, d1, d2)

	assert.Equal(t, "merge_request_hooks", *got2)
	assert.Equal(t, "fasgasd", *got3)
	assert.Equal(t, false, got4)

	pr, _ := got1.(*PullRequestEvent)
	assert.Equal(t, "open", *pr.GetAction())
	assert.Equal(t, "opened", *pr.GetState())
	assert.Equal(t, "ibforuorg", *pr.GetOrg())
	assert.Equal(t, "test1", *pr.GetRepo())
	assert.Equal(t, "https://gitcode.com/ibforuorg/test1", *pr.GetHtmlURL())
	assert.Equal(t, (*string)(nil), pr.GetBase())
	assert.Equal(t, (*string)(nil), pr.GetHead())
	assert.Equal(t, "4", *pr.GetNumber())
	assert.Equal(t, "ibforu", *pr.GetAuthor())
	assert.Equal(t, (*string)(nil), pr.GetComment())
	assert.Equal(t, (*string)(nil), pr.GetCommenter())
	assert.Equal(t, 0, len(pr.ListLabels()))

	pr = new(PullRequestEvent)
	assert.Equal(t, (*string)(nil), pr.GetAction())
	assert.Equal(t, (*string)(nil), pr.GetState())
	assert.Equal(t, (*string)(nil), pr.GetOrg())
	assert.Equal(t, (*string)(nil), pr.GetRepo())
	assert.Equal(t, (*string)(nil), pr.GetHtmlURL())
	assert.Equal(t, (*string)(nil), pr.GetBase())
	assert.Equal(t, (*string)(nil), pr.GetHead())
	assert.Equal(t, (*string)(nil), pr.GetNumber())
	assert.Equal(t, (*string)(nil), pr.GetAuthor())
	assert.Equal(t, (*string)(nil), pr.GetComment())
	assert.Equal(t, (*string)(nil), pr.GetCommenter())
	assert.Equal(t, 0, len(pr.ListLabels()))
}

func notePR(t *testing.T) {
	want := GitCodeAccessor{Note: new(NoteEvent)}
	data := readWebHookTestdata(t, webhookTestDataDir+"pr_note.json", want.Note)

	buf := &bytes.Buffer{}
	buf.Write(data)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/2", buf)
	req.Header.Set(headerEventType, "note_hooks")
	req.Header.Set(headerEventGUID, "651234123")
	w := httptest.NewRecorder()

	a := new(GitCodeAccessor)
	got1, got2, got3, got4 := a.GetAccessor(w, req)
	d1, _ := json.Marshal(want.Note)
	d2, _ := json.Marshal(got1)
	assert.Equal(t, d1, d2)

	assert.Equal(t, "note_hooks", *got2)
	assert.Equal(t, "651234123", *got3)
	assert.Equal(t, true, got4)

	note, _ := got1.(*NoteEvent)
	assert.Equal(t, "open", *note.GetAction())
	assert.Equal(t, "opened", *note.GetState())
	assert.Equal(t, "ibforuorg", *note.GetOrg())
	assert.Equal(t, "test1", *note.GetRepo())
	assert.Equal(t, "https://gitcode.com/ibforuorg/test1", *note.GetHtmlURL())
	assert.Equal(t, (*string)(nil), note.GetBase())
	assert.Equal(t, (*string)(nil), note.GetHead())
	assert.Equal(t, "4", *note.GetNumber())
	assert.Equal(t, "ibforu", *note.GetAuthor())
	assert.Equal(t, "/lgtm\n/approve", *note.GetComment())
	assert.Equal(t, "ibforu2nd", *note.GetCommenter())
	assert.Equal(t, 0, len(note.ListLabels()))
}

func noteIssue(t *testing.T) {
	want := GitCodeAccessor{Note: new(NoteEvent)}
	data := readWebHookTestdata(t, webhookTestDataDir+"issues_note.json", want.Note)

	buf := &bytes.Buffer{}
	buf.Write(data)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/2", buf)
	req.Header.Set(headerEventType, "note_hooks")
	req.Header.Set(headerEventGUID, "151231321")
	w := httptest.NewRecorder()

	a := new(GitCodeAccessor)
	got1, got2, got3, got4 := a.GetAccessor(w, req)
	d1, _ := json.Marshal(want.Note)
	d2, _ := json.Marshal(got1)
	assert.Equal(t, d1, d2)

	assert.Equal(t, "note_hooks", *got2)
	assert.Equal(t, "151231321", *got3)
	assert.Equal(t, false, got4)

	note, _ := got1.(*NoteEvent)
	assert.Equal(t, (*string)(nil), note.GetAction())
	assert.Equal(t, "opened", *note.GetState())
	assert.Equal(t, "ibforuorg", *note.GetOrg())
	assert.Equal(t, "test1", *note.GetRepo())
	assert.Equal(t, "https://gitcode.com/ibforuorg/test1", *note.GetHtmlURL())
	assert.Equal(t, (*string)(nil), note.GetBase())
	assert.Equal(t, (*string)(nil), note.GetHead())
	assert.Equal(t, "4", *note.GetNumber())
	assert.Equal(t, "ibforu", *note.GetAuthor())
	assert.Equal(t, "oiugbfaijub", *note.GetComment())
	assert.Equal(t, "ibforu2nd", *note.GetCommenter())
	assert.Equal(t, 0, len(note.ListLabels()))

	note = new(NoteEvent)
	assert.Equal(t, (*string)(nil), note.GetAction())
	assert.Equal(t, (*string)(nil), note.GetState())
	assert.Equal(t, (*string)(nil), note.GetOrg())
	assert.Equal(t, (*string)(nil), note.GetRepo())
	assert.Equal(t, (*string)(nil), note.GetHtmlURL())
	assert.Equal(t, (*string)(nil), note.GetBase())
	assert.Equal(t, (*string)(nil), note.GetHead())
	assert.Equal(t, (*string)(nil), note.GetNumber())
	assert.Equal(t, (*string)(nil), note.GetAuthor())
	assert.Equal(t, (*string)(nil), note.GetComment())
	assert.Equal(t, (*string)(nil), note.GetCommenter())
	assert.Equal(t, 0, len(note.ListLabels()))
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
