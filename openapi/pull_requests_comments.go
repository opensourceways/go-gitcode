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

// 获取某个Pull Request的所有评论

// CreatePullRequestComment 提交pull request 评论
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#4-%e6%8f%90%e4%ba%a4pull-request-%e8%af%84%e8%ae%ba
func (s *PullRequestsService) CreatePullRequestComment(ctx context.Context, owner, repo, number string, comment *PullRequestComment) (*PullRequestComment, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/comments", owner, repo, number)
	req, err := s.api.newRequest(http.MethodPost, urlStr, comment)
	if err != nil {
		return nil, false, err
	}

	addedComment := new(PullRequestComment)
	resp, err := s.api.Do(ctx, req, addedComment)
	return addedComment, successCreated(resp), err
}

// EditPullRequestComment 编辑评论
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#27-%e7%bc%96%e8%be%91%e8%af%84%e8%ae%ba
func (s *PullRequestsService) EditPullRequestComment(ctx context.Context, owner, repo, number string, comment *PullRequestComment) (*PullRequestComment, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/comments", owner, repo, number)
	req, err := s.api.newRequest(http.MethodPost, urlStr, comment)
	if err != nil {
		return nil, false, err
	}

	addedComment := new(PullRequestComment)
	resp, err := s.api.Do(ctx, req, addedComment)
	return addedComment, successModified(resp), err
}

// DeletePullRequestComment 删除评论
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#28-%e5%88%a0%e9%99%a4%e8%af%84%e8%ae%ba
func (s *PullRequestsService) DeletePullRequestComment(ctx context.Context, owner, repo, number, commentId string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/comments/%s", owner, repo, number, commentId)
	req, err := s.api.newRequest(http.MethodDelete, urlStr, commentId)
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}
