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
	"github.com/opensourceways/go-gitcode/openapi"
	"strconv"
)

type Project struct {
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	HTMLURL   *string `json:"web_url,omitempty"`
}

type Attributes struct {
	Action          *string `json:"action,omitempty"`
	State           *string `json:"state,omitempty"`
	Number          *int    `json:"iid,omitempty"`
	CommentID       *string `json:"discussion_id,omitempty"`
	Comment         *string `json:"note,omitempty"`
	CommentCategory *string `json:"noteable_type,omitempty"`
}

type IssuePart struct {
	Action *string       `json:"action,omitempty"`
	State  *string       `json:"state,omitempty"`
	Number *int          `json:"iid,omitempty"`
	Author *openapi.User `json:"author,omitempty"`
}

type IssueEvent struct {
	UUID        *string          `json:"uuid,omitempty"`
	EventType   *string          `json:"event_type,omitempty"`
	ObjectKind  *string          `json:"object_kind,omitempty"`
	ManualBuild *bool            `json:"manual_build,omitempty"`
	Attributes  *Attributes      `json:"object_attributes,omitempty"`
	User        *openapi.User    `json:"user,omitempty"`
	Assignees   []*openapi.User  `json:"assignees,omitempty"`
	Repository  *Project         `json:"project,omitempty"`
	Labels      []*openapi.Label `json:"labels,omitempty"`
	Issue       *IssuePart       `json:"issue,omitempty"`
}

func (iss *IssueEvent) GetAction() *string {
	if iss.Attributes == nil {
		return nil
	}

	return iss.Attributes.Action
}
func (iss *IssueEvent) GetState() *string {
	if iss.Attributes == nil {
		return nil
	}

	return iss.Attributes.State
}
func (iss *IssueEvent) GetOrg() *string {
	if iss.Repository == nil {
		return nil
	}

	return iss.Repository.Namespace
}
func (iss *IssueEvent) GetRepo() *string {
	if iss.Repository == nil {
		return nil
	}

	return iss.Repository.Name
}
func (iss *IssueEvent) GetHtmlURL() *string {
	if iss.Repository == nil {
		return nil
	}

	return iss.Repository.HTMLURL
}
func (iss *IssueEvent) GetBase() *string {
	return nil
}
func (iss *IssueEvent) GetHead() *string {
	return nil
}
func (iss *IssueEvent) GetNumber() *string {
	if iss.Attributes != nil && iss.Attributes.Number != nil {
		n := strconv.Itoa(*iss.Attributes.Number)
		return &n
	}

	if iss.Issue != nil && iss.Issue.Number != nil {
		n := strconv.Itoa(*iss.Issue.Number)
		return &n
	}

	return nil
}
func (iss *IssueEvent) GetAuthor() *string {
	if iss.User == nil {
		return nil
	}
	if iss.User.Login != nil {
		return iss.User.Login
	}

	return iss.User.UserName
}
func (iss *IssueEvent) GetComment() *string {
	return nil
}
func (iss *IssueEvent) GetCommenter() *string {
	return nil
}
func (iss *IssueEvent) ListLabels() []*string {
	if len(iss.Labels) == 0 {
		return nil
	}

	labels := make([]*string, 0, len(iss.Labels))
	for _, p := range iss.Labels {
		labels = append(labels, &p.Name)
	}
	return labels
}
