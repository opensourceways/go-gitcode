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
)

// GetPullRequest 获取单个Pull Requests
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#10-%e8%8e%b7%e5%8f%96%e5%8d%95%e4%b8%aapull-request
func (s *PullRequestsService) GetPullRequest(ctx context.Context, owner, repo, number string) (*PullRequest, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s", owner, repo, number)
	req, err := s.api.newRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	pr := new(PullRequest)
	resp, err := s.api.Do(ctx, req, pr)
	return pr, successGetData(resp), err
}

// 创建Pull Request

// UpdatePullRequest 更新Pull Request信息
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#8-%e6%9b%b4%e6%96%b0pull-request%e4%bf%a1%e6%81%af
func (s *PullRequestsService) UpdatePullRequest(ctx context.Context, owner, repo, number string, prContent *PullRequestRequest) (*PullRequest, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s", owner, repo, number)
	req, err := s.api.newRequest(http.MethodPost, urlStr, prContent)
	if err != nil {
		return nil, false, err
	}

	pr := new(PullRequest)
	resp, err := s.api.Do(ctx, req, pr)
	return pr, successCreated(resp), err
}

// ListPullRequestOperationLog 获取某个Pull Request的操作日志
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#16-%e8%8e%b7%e5%8f%96%e6%9f%90%e4%b8%aapull-request%e7%9a%84%e6%93%8d%e4%bd%9c%e6%97%a5%e5%bf%97
func (s *PullRequestsService) ListPullRequestOperationLog(ctx context.Context, owner, repo, number string) ([]*PullRequestOperationLog, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/operate_logs", owner, repo, number)
	req, err := s.api.newRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	var operationLogs []*PullRequestOperationLog
	resp, err := s.api.Do(ctx, req, &operationLogs)
	return operationLogs, successGetData(resp), err
}

// ListPullRequestCommits 获取某Pull Request的所有Commit信息
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#11-%e8%8e%b7%e5%8f%96%e6%9f%90pull-request%e7%9a%84%e6%89%80%e6%9c%89commit%e4%bf%a1%e6%81%af
func (s *PullRequestsService) ListPullRequestCommits(ctx context.Context, owner, repo, number string) ([]*RepositoryCommit, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/commits", owner, repo, number)
	req, err := s.api.newRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	var prCommits []*RepositoryCommit
	resp, err := s.api.Do(ctx, req, &prCommits)
	return prCommits, successGetData(resp), err
}

// ListPullRequestFiles Pull Request Commit文件列表
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#5-pull-request-commit%e6%96%87%e4%bb%b6%e5%88%97%e8%a1%a8
func (s *PullRequestsService) ListPullRequestFiles(ctx context.Context, owner, repo, number string) ([]*CommitFile, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/files", owner, repo, number)
	req, err := s.api.newRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	var prFiles []*CommitFile
	resp, err := s.api.Do(ctx, req, &prFiles)
	return prFiles, successGetData(resp), err
}

// MergePullRequest 合并Pull Request
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#2-%e5%90%88%e5%b9%b6pull-request
func (s *PullRequestsService) MergePullRequest(ctx context.Context, owner, repo, number string) (*PullRequestMergedResult, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/merge", owner, repo, number)
	req, err := s.api.newRequest(http.MethodPut, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	mergeResult := new(PullRequestMergedResult)
	resp, err := s.api.Do(ctx, req, mergeResult)
	return mergeResult, successModified(resp), err
}

// ListPullRequestLinkingIssues 获取pr关联的issue
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#3-%e8%8e%b7%e5%8f%96pr%e5%85%b3%e8%81%94%e7%9a%84issue
func (s *PullRequestsService) ListPullRequestLinkingIssues(ctx context.Context, owner, repo, number string) ([]*Issue, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pull_requests/%s/issues", owner, repo, number)
	req, err := s.api.newRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	var linkingPRList []*Issue
	resp, err := s.api.Do(ctx, req, &linkingPRList)
	return linkingPRList, successGetData(resp), err
}
