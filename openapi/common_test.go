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
	"errors"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewRequestError(t *testing.T) {

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

	issue, ok, err := client.Issues.UpdateIssue(context.Background(), owner, "2", &IssueRequest{
		Repository: repo,
		Title:      "issue1",
	})
	assert.Equal(t, false, ok)
	assert.Equal(t, (*Issue)(nil), issue)
	assert.Equal(t, msg1, err.Error())

	pr, ok, err := client.Issues.ListIssueLinkingPullRequests(context.Background(), owner, repo, "1")
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*PullRequest)(nil), pr)
	assert.Equal(t, msg1, err.Error())

	comment := "fajhgdahjksghj"
	issueComment, ok, err := client.Issues.CreateIssueComment(context.Background(), owner, repo, "1", &IssueComment{
		Body: &comment,
	})
	assert.Equal(t, false, ok)
	assert.Equal(t, (*IssueComment)(nil), issueComment)
	assert.Equal(t, msg1, err.Error())

	labels, ok, err := client.Issues.ListRepoIssueLabels(context.Background(), owner, repo)
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*Label)(nil), labels)
	assert.Equal(t, msg1, err.Error())
	result1, ok, err := client.Issues.CreateRepoIssueLabel(context.Background(), owner, repo, &Label{Name: "fasdsad", Color: "#fff"})
	assert.Equal(t, false, ok)
	assert.Equal(t, (*Label)(nil), result1)
	assert.Equal(t, msg1, err.Error())
	result2, ok, err := client.Issues.UpdateRepoIssueLabel(context.Background(), owner, repo, "giasdlkggds", "fsaghhhhh", "#000000")
	assert.Equal(t, false, ok)
	assert.Equal(t, (*Label)(nil), result2)
	assert.Equal(t, msg1, err.Error())
	ok, err = client.Issues.DeleteRepoIssueLabel(context.Background(), owner, repo, "fgagasdasda")
	assert.Equal(t, false, ok)
	assert.Equal(t, msg1, err.Error())
	result3, ok, err := client.Issues.AddLabelsToIssue(context.Background(), owner, repo, "1", []string{"fsaghhhhh"})
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*Label)(nil), result3)
	assert.Equal(t, msg1, err.Error())
	ok, err = client.Issues.RemoveLabelsFromIssue(context.Background(), owner, repo, "1", "fsaghhhhh")
	assert.Equal(t, false, ok)
	assert.Equal(t, msg1, err.Error())

	result4, ok, err := client.PullRequests.GetPullRequest(context.Background(), owner, repo, "1")
	assert.Equal(t, false, ok)
	assert.Equal(t, (*PullRequest)(nil), result4)
	assert.Equal(t, msg1, err.Error())
	result5, ok, err := client.PullRequests.UpdatePullRequest(context.Background(), owner, repo, "1", &PullRequestRequest{
		State: "open",
	})
	assert.Equal(t, false, ok)
	assert.Equal(t, (*PullRequest)(nil), result5)
	assert.Equal(t, msg1, err.Error())
	result6, ok, err := client.PullRequests.ListPullRequestLinkingIssues(context.Background(), owner, repo, "1")
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*Issue)(nil), result6)
	assert.Equal(t, msg1, err.Error())
	result7, ok, err := client.PullRequests.CreatePullRequestComment(context.Background(), owner, repo, "1", &PullRequestCommentRequest{
		Body: "fauygiahsgdbviahsd",
	})
	assert.Equal(t, false, ok)
	assert.Equal(t, (*SimpleComment)(nil), result7)
	assert.Equal(t, msg1, err.Error())
	result8, ok, err := client.PullRequests.AddLabelsToPullRequest(context.Background(), owner, repo, "1", []string{"bug1"})
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*Label)(nil), result8)
	assert.Equal(t, msg1, err.Error())
	ok, err = client.PullRequests.RemoveLabelsFromPullRequest(context.Background(), owner, repo, "1", []string{"bug1"})
	result9, ok, err := client.Repository.CheckUserIsRepoMember(context.Background(), owner, repo, "ibforu2nd")
	assert.Equal(t, false, ok)
	assert.Equal(t, false, result9)
	assert.Equal(t, msg1, err.Error())

	result10, ok, err := client.Repository.GetRepoAllMember(context.Background(), owner, repo, "1")
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*User)(nil), result10)
	assert.Equal(t, msg1, err.Error())

	result11, okAndNg, err := client.Repository.GetRepoMemberPermission(context.Background(), owner, repo, "1")
	assert.Equal(t, [2]bool{false, false}, okAndNg)
	assert.Equal(t, (*User)(nil), result11)
	assert.Equal(t, msg1, err.Error())

	result12, ok, err := client.Repository.GetRepoContributors(context.Background(), owner, repo, "")
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*Contributor)(nil), result12)
	assert.Equal(t, msg1, err.Error())

	result13, ok, err := client.User.GetUserInfo(context.Background())
	assert.Equal(t, false, ok)
	assert.Equal(t, (*User)(nil), result13)
	assert.Equal(t, msg1, err.Error())

	result14, ok, err := client.PullRequests.ListPullRequestCommits(context.Background(), owner, repo, "1")
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*RepositoryCommit)(nil), result14)
	assert.Equal(t, msg1, err.Error())

	result15, ok, err := client.Issues.GetIssueLabels(context.Background(), owner, "15423", "1")
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*Label)(nil), result15)
	assert.Equal(t, msg1, err.Error())

	result16, ok, err := client.PullRequests.GetPullRequestChangeFiles(context.Background(), owner, repo, "1")
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*CommitFile)(nil), result16)
	assert.Equal(t, msg1, err.Error())

	result17, ok, err := client.PullRequests.ListPullRequestOperationLogs(context.Background(), owner, repo, "1", "asc", "1")
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*PullRequestOperationLog)(nil), result17)
	assert.Equal(t, msg1, err.Error())

	result18, ok, err := client.PullRequests.MergePullRequest(context.Background(), owner, repo, "1", "merge")
	assert.Equal(t, false, ok)
	assert.Equal(t, (*PullRequestMergedResult)(nil), result18)
	assert.Equal(t, msg1, err.Error())

	result19, ok, err := client.PullRequests.ListPullRequestComments(context.Background(), owner, repo, "1", "1", "pr_comment")
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*PullRequestComment)(nil), result19)
	assert.Equal(t, msg1, err.Error())

	ok, err = client.PullRequests.UpdatePullRequestComment(context.Background(), owner, repo, "11234", "1fasd")
	assert.Equal(t, false, ok)
	assert.Equal(t, msg1, err.Error())

	ok, err = client.PullRequests.DeletePullRequestComment(context.Background(), owner, repo, "11234")
	assert.Equal(t, false, ok)
	assert.Equal(t, msg1, err.Error())

	result20, ok, err := client.PullRequests.GetLabelsOfPullRequest(context.Background(), owner, repo, "1")
	assert.Equal(t, false, ok)
	assert.Equal(t, ([]*Label)(nil), result20)
	assert.Equal(t, msg1, err.Error())

	result21, ok, err := client.Repository.GetRepoContentByPath(context.Background(), owner, repo, "1.txt", "main")
	assert.Equal(t, false, ok)
	assert.Equal(t, (*RepositoryContent)(nil), result21)
	assert.Equal(t, msg1, err.Error())
}
