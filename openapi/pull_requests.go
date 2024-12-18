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

// GetPullRequest 获取单个Pull Requests
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#10-%e8%8e%b7%e5%8f%96%e5%8d%95%e4%b8%aapull-request
func (s *PullRequestsService) GetPullRequest(ctx context.Context, owner, repo, number string) (*PullRequest, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s", owner, repo, number)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	pr := new(PullRequest)
	resp, err := s.api.Do(ctx, req, pr)
	return pr, successGetData(resp), err
}

// UpdatePullRequest 更新Pull Request信息
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#8-%e6%9b%b4%e6%96%b0pull-request%e4%bf%a1%e6%81%af
func (s *PullRequestsService) UpdatePullRequest(ctx context.Context, owner, repo, number string, prContent *PullRequestRequest) (*PullRequest, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s", owner, repo, number)
	req, err := newRequest(s.api, http.MethodPatch, urlStr, prContent)
	if err != nil {
		return nil, false, err
	}

	pr := new(PullRequest)
	resp, err := s.api.Do(ctx, req, pr)
	return pr, successCreated(resp), err
}

// ListPullRequestLinkingIssues 获取pr关联的issue
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#3-%e8%8e%b7%e5%8f%96pr%e5%85%b3%e8%81%94%e7%9a%84issue
func (s *PullRequestsService) ListPullRequestLinkingIssues(ctx context.Context, owner, repo, number string) ([]*Issue, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/issues", owner, repo, number)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	var linkingPRList []*Issue
	resp, err := s.api.Do(ctx, req, &linkingPRList)
	return linkingPRList, successGetData(resp), err
}

// ListPullRequestCommits 获取某Pull Request的所有Commit信息
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#10-%e8%8e%b7%e5%8f%96%e6%9f%90pull-request%e7%9a%84%e6%89%80%e6%9c%89commit%e4%bf%a1%e6%81%af
func (s *PullRequestsService) ListPullRequestCommits(ctx context.Context, owner, repo, number string) ([]*RepositoryCommit, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/commits", owner, repo, number)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	var commitList []*RepositoryCommit
	resp, err := s.api.Do(ctx, req, &commitList)
	return commitList, successGetData(resp), err
}

// GetPullRequestChangeFiles Pull Request Commit文件列表
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#5-pull-request-commit%e6%96%87%e4%bb%b6%e5%88%97%e8%a1%a8
func (s *PullRequestsService) GetPullRequestChangeFiles(ctx context.Context, owner, repo, number string) ([]*CommitFile, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/files", owner, repo, number)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	var changes []*CommitFile
	resp, err := s.api.Do(ctx, req, &changes)
	return changes, successGetData(resp), err
}

// MergePullRequest 合并Pull Request
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#2-%e5%90%88%e5%b9%b6pull-request
func (s *PullRequestsService) MergePullRequest(ctx context.Context, owner, repo, number, mergeMethod string) (*PullRequestMergedResult, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/merge", owner, repo, number)
	req, err := newRequest(s.api, http.MethodPut, urlStr, &PullRequestRequestMerge{
		Method: mergeMethod,
	})
	if err != nil {
		return nil, false, err
	}

	pr := new(PullRequestMergedResult)
	resp, err := s.api.Do(ctx, req, pr)
	return pr, successCreated(resp), err
}

// ListPullRequestOperationLogs 获取某个Pull Request的操作日志
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#15-%e8%8e%b7%e5%8f%96%e6%9f%90%e4%b8%aapull-request%e7%9a%84%e6%93%8d%e4%bd%9c%e6%97%a5%e5%bf%97
func (s *PullRequestsService) ListPullRequestOperationLogs(ctx context.Context, owner, repo, number, sort, page string) ([]*PullRequestOperationLog, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/operate_logs", owner, repo, number)
	req, err := newRequest(s.api, http.MethodGet, urlStr,
		&url.Values{"page": []string{page}, "per_page": []string{"100"}, "sort": []string{sort}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var logs []*PullRequestOperationLog
	resp, err := s.api.Do(ctx, req, &logs)
	return logs, successGetData(resp), err
}
