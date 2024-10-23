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
	"fmt"
	"net/http"
	"net/url"
)

// UpdateIssue 更新Issue
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/issues/#2-%e6%9b%b4%e6%96%b0issue
func (s *IssuesService) UpdateIssue(ctx context.Context, owner, number string, issueContent *IssueRequest) (*Issue, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/issues/%s", owner, number)
	req, err := s.api.newRequest(http.MethodPatch, urlStr, issueContent)
	if err != nil {
		return nil, false, err
	}

	updatedIssue := new(Issue)
	resp, err := s.api.Do(ctx, req, updatedIssue)
	return updatedIssue, successCreated(resp), err
}

// ListIssueLinkingPullRequests 获取 issue 关联的 pull requests
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/issues/#7-%e8%8e%b7%e5%8f%96-issue-%e5%85%b3%e8%81%94%e7%9a%84-pull-requests
func (s *IssuesService) ListIssueLinkingPullRequests(ctx context.Context, owner, repo, number string) ([]*PullRequest, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/issues/%s/pull_requests", owner, number)
	req, err := s.api.newRequest(http.MethodGet, urlStr, &url.Values{"repo": []string{repo}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var linkingPRList []*PullRequest
	resp, err := s.api.Do(ctx, req, &linkingPRList)
	return linkingPRList, successGetData(resp), err
}
