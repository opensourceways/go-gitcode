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
	"encoding/json"
	"github.com/opensourceways/go-gitcode/openapi"
	"strconv"
)

type PRPart struct {
	Action       *string       `json:"action,omitempty"`
	State        *string       `json:"state,omitempty"`
	Number       *int          `json:"iid,omitempty"`
	Author       *openapi.User `json:"author,omitempty"`
	TargetBranch *string       `json:"target_branch,omitempty"`
	Source       *Project      `json:"source,omitempty"`
	SourceBranch *string       `json:"source_branch,omitempty"`
	ID           *json.Number  `json:"id,omitempty"`
}

type PullRequestEvent struct {
	UUID        *string          `json:"uuid,omitempty"`
	EventType   *string          `json:"event_type,omitempty"`
	ObjectKind  *string          `json:"object_kind,omitempty"`
	ManualBuild *bool            `json:"manual_build,omitempty"`
	Attributes  *Attributes      `json:"object_attributes,omitempty"`
	User        *openapi.User    `json:"user,omitempty"`
	Repository  *Project         `json:"project,omitempty"`
	Labels      []*openapi.Label `json:"labels,omitempty"`
	PR          *PRPart          `json:"merge_request,omitempty"`
}

func (pr *PullRequestEvent) GetAction() *string {
	if pr.Attributes == nil {
		return nil
	}

	return pr.Attributes.Action
}
func (pr *PullRequestEvent) GetActionDetail() *string {
	if pr.Attributes == nil {
		return nil
	}

	return pr.Attributes.ActionDetail
}
func (pr *PullRequestEvent) GetState() *string {
	if pr.Attributes == nil {
		return nil
	}

	return pr.Attributes.State
}
func (pr *PullRequestEvent) GetOrg() *string {
	if pr.Repository == nil {
		return nil
	}

	return pr.Repository.Namespace
}
func (pr *PullRequestEvent) GetRepo() *string {
	if pr.Repository == nil {
		return nil
	}

	return pr.Repository.Name
}
func (pr *PullRequestEvent) GetHtmlURL() *string {
	if pr.Attributes == nil {
		return nil
	}

	return pr.Attributes.URL
}
func (pr *PullRequestEvent) GetBase() *string {
	if pr.Attributes == nil {
		return nil
	}

	return pr.Attributes.TargetBranch
}
func (pr *PullRequestEvent) GetHead() *string {
	if pr.Attributes == nil || pr.Attributes.SourceBranch == nil ||
		pr.Attributes.Source == nil || pr.Attributes.Source.Path == nil {
		return nil
	}

	head := *pr.Attributes.Source.Path + "/" + *pr.Attributes.SourceBranch
	return &head
}
func (pr *PullRequestEvent) GetNumber() *string {
	if pr.Attributes != nil && pr.Attributes.Number != nil {
		n := strconv.Itoa(*pr.Attributes.Number)
		return &n
	}

	return nil
}
func (pr *PullRequestEvent) GetID() *string {
	if pr.Attributes != nil && pr.Attributes.ID != nil {
		n := pr.Attributes.ID.String()
		return &n
	}

	return nil
}
func (pr *PullRequestEvent) GetAuthor() *string {
	if pr.User == nil {
		return nil
	}

	return pr.User.UserName
}
func (pr *PullRequestEvent) GetCommentID() *string {
	return nil
}
func (pr *PullRequestEvent) GetCommentKind() *string {
	return nil
}
func (pr *PullRequestEvent) GetComment() *string {
	return nil
}
func (pr *PullRequestEvent) GetCommenter() *string {
	return nil
}
func (pr *PullRequestEvent) GetCreateTime() *string {
	if pr.Attributes == nil || pr.Attributes.CreateTime == nil {
		return nil
	}

	return pr.Attributes.CreateTime.ToString()
}

func (pr *PullRequestEvent) GetUpdateTime() *string {
	if pr.Attributes == nil || pr.Attributes.UpdatedTime == nil {
		return nil
	}

	return pr.Attributes.UpdatedTime.ToString()
}
