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

// CreatePullRequestComment 提交 pull request 评论
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#4-%e6%8f%90%e4%ba%a4pull-request-%e8%af%84%e8%ae%ba
func (s *PullRequestsService) CreatePullRequestComment(ctx context.Context, owner, repo, number string, comment *PullRequestCommentRequest) (*PullRequestComment, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/comments", owner, repo, number)
	req, err := newRequest(s.api, http.MethodPost, urlStr, comment)
	if err != nil {
		return nil, false, err
	}

	addedComment := new(PullRequestComment)
	resp, err := s.api.Do(ctx, req, addedComment)
	return addedComment, successCreated(resp), err
}

// ListPullRequestComments 获取某个Pull Request的所有评论
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#6-%e8%8e%b7%e5%8f%96%e6%9f%90%e4%b8%aapull-request%e7%9a%84%e6%89%80%e6%9c%89%e8%af%84%e8%ae%ba
func (s *PullRequestsService) ListPullRequestComments(ctx context.Context, owner, repo, number, page, commentType string) ([]*PullRequestComment, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/comments", owner, repo, number)
	req, err := newRequest(s.api, http.MethodGet, urlStr,
		&url.Values{"page": []string{page}, "per_page": []string{"100"}, "comment_type": []string{commentType}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var comments []*PullRequestComment
	resp, err := s.api.Do(ctx, req, &comments)
	return comments, successGetData(resp), err
}

// UpdatePullRequestComment 编辑评论
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#26-%e7%bc%96%e8%be%91%e8%af%84%e8%ae%ba
func (s *PullRequestsService) UpdatePullRequestComment(ctx context.Context, owner, repo, commentID, body string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/comments/%s", owner, repo, commentID)
	req, err := newRequest(s.api, http.MethodPatch, urlStr, &PullRequestComment{
		Body: &body,
	})
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// DeletePullRequestComment 删除评论
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#27-%e5%88%a0%e9%99%a4%e8%af%84%e8%ae%ba
func (s *PullRequestsService) DeletePullRequestComment(ctx context.Context, owner, repo, commentID string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/comments/%s", owner, repo, commentID)
	req, err := newRequest(s.api, http.MethodDelete, urlStr, nil)
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}
