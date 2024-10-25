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
	"errors"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewRequestError(t *testing.T) {
	t.Parallel()
	client, _, _ := mockServer(t)

	msg1 := "1231412"
	patch := gomonkey.ApplyFunc(newRequest, func(api *APIClient, method, urlStr string, body any, handlers ...RequestHandler) (*http.Request, error) {
		return nil, errors.New(msg1)
	})

	defer patch.Reset()
	targetLabels, ok, err := client.Issues.ListRepoIssueLabels(context.Background(), owner, repo)
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, len(targetLabels))
	assert.Equal(t, msg1, err.Error())

	issue, ok, err := client.Issues.UpdateIssue(context.Background(), "ibforuorg", "2", &IssueRequest{
		Repository: "test1",
		Title:      "issue1",
	})
	assert.Equal(t, false, ok)
	assert.Equal(t, (*Issue)(nil), issue)
	assert.Equal(t, msg1, err.Error())

	pr, ok, err := client.Issues.ListIssueLinkingPullRequests(context.Background(), "ibforuorg", "test1", "1")
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*PullRequest)(nil), pr)
	assert.Equal(t, msg1, err.Error())

	comment := "fajhgdahjksghj"
	issueComment, ok, err := client.Issues.CreateIssueComment(context.Background(), "ibforuorg", "test1", "1", &IssueComment{
		Body: &comment,
	})
	assert.Equal(t, false, ok)
	assert.Equal(t, (*IssueComment)(nil), issueComment)
	assert.Equal(t, msg1, err.Error())

	labels, ok, err := client.Issues.ListRepoIssueLabels(context.Background(), "ibforuorg", "test1")
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*Label)(nil), labels)
	assert.Equal(t, msg1, err.Error())
	result1, ok, err := client.Issues.CreateRepoIssueLabel(context.Background(), "ibforuorg", "test1", &Label{Name: "fasdsad", Color: "#fff"})
	assert.Equal(t, false, ok)
	assert.Equal(t, (*Label)(nil), result1)
	assert.Equal(t, msg1, err.Error())
	result2, ok, err := client.Issues.UpdateRepoIssueLabel(context.Background(), "ibforuorg", "test1", "giasdlkggds", "fsaghhhhh", "#000000")
	assert.Equal(t, false, ok)
	assert.Equal(t, (*Label)(nil), result2)
	assert.Equal(t, msg1, err.Error())
	ok, err = client.Issues.DeleteRepoIssueLabel(context.Background(), "ibforuorg", "test1", "fgagasdasda")
	assert.Equal(t, false, ok)
	assert.Equal(t, msg1, err.Error())
	result3, ok, err := client.Issues.AddLabelsToIssue(context.Background(), "ibforuorg", "test1", "1", []string{"fsaghhhhh"})
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*Label)(nil), result3)
	assert.Equal(t, msg1, err.Error())
	ok, err = client.Issues.RemoveLabelsFromIssue(context.Background(), "ibforuorg", "test1", "1", "fsaghhhhh")
	assert.Equal(t, false, ok)
	assert.Equal(t, msg1, err.Error())

	result4, ok, err := client.PullRequests.GetPullRequest(context.Background(), "ibforuorg", "test1", "1")
	assert.Equal(t, false, ok)
	assert.Equal(t, (*PullRequest)(nil), result4)
	assert.Equal(t, msg1, err.Error())
	result5, ok, err := client.PullRequests.UpdatePullRequest(context.Background(), "ibforuorg", "test1", "1", &PullRequestRequest{
		State: "open",
	})
	assert.Equal(t, false, ok)
	assert.Equal(t, (*PullRequest)(nil), result5)
	assert.Equal(t, msg1, err.Error())
	result6, ok, err := client.PullRequests.ListPullRequestLinkingIssues(context.Background(), "ibforuorg", "test1", "1")
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*Issue)(nil), result6)
	assert.Equal(t, msg1, err.Error())
	result7, ok, err := client.PullRequests.CreatePullRequestComment(context.Background(), "ibforuorg", "test1", "1", &PullRequestCommentRequest{
		Body: "fauygiahsgdbviahsd",
	})
	assert.Equal(t, false, ok)
	assert.Equal(t, (*PullRequestComment)(nil), result7)
	assert.Equal(t, msg1, err.Error())
	result8, ok, err := client.PullRequests.AddLabelsToPullRequest(context.Background(), "ibforuorg", "test1", "1", []string{"bug1"})
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*Label)(nil), result8)
	assert.Equal(t, msg1, err.Error())
	ok, err = client.PullRequests.RemoveLabelsFromPullRequest(context.Background(), "ibforuorg", "test1", "1", "bug1")
	result9, ok, err := client.Repository.CheckUserIsRepoMember(context.Background(), "ibforuorg", "test1", "ibforu2nd")
	assert.Equal(t, false, ok)
	assert.Equal(t, false, result9)
	assert.Equal(t, msg1, err.Error())

	result10, ok, err := client.Repository.GetRepoAllMember(context.Background(), "ibforuorg", "test1", "1")
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*User)(nil), result10)
	assert.Equal(t, msg1, err.Error())

	result11, ok, err := client.Repository.GetRepoMemberPermission(context.Background(), "ibforuorg", "test1", "1")
	assert.Equal(t, false, ok)
	assert.Equal(t, (*User)(nil), result11)
	assert.Equal(t, msg1, err.Error())
}
