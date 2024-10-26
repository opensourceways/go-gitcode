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
	"github.com/opensourceways/go-gitcode/openapi"
	"strconv"
)

type NoteEvent struct {
	UUID        *string          `json:"uuid,omitempty"`
	EventType   *string          `json:"event_type,omitempty"`
	ObjectKind  *string          `json:"object_kind,omitempty"`
	ManualBuild *bool            `json:"manual_build,omitempty"`
	Attributes  *Attributes      `json:"object_attributes,omitempty"`
	User        *openapi.User    `json:"user,omitempty"`
	Repository  *Project         `json:"project,omitempty"`
	Labels      []*openapi.Label `json:"labels,omitempty"`
	Issue       *IssuePart       `json:"issue,omitempty"`
	PR          *PRPart          `json:"merge_request,omitempty"`
}

func (n *NoteEvent) GetAction() *string {
	if n.Issue != nil && n.Issue.Action != nil {
		return n.Issue.Action
	}

	if n.PR != nil && n.PR.Action != nil {
		return n.PR.Action
	}

	return nil
}
func (n *NoteEvent) GetState() *string {
	if n.Issue != nil && n.Issue.State != nil {
		return n.Issue.State
	}

	if n.PR != nil && n.PR.State != nil {
		return n.PR.State
	}

	return nil
}
func (n *NoteEvent) GetOrg() *string {
	if n.Repository == nil {
		return nil
	}

	return n.Repository.Namespace
}
func (n *NoteEvent) GetRepo() *string {
	if n.Repository == nil {
		return nil
	}

	return n.Repository.Name
}
func (n *NoteEvent) GetHtmlURL() *string {
	if n.Repository == nil {
		return nil
	}

	return n.Repository.HTMLURL
}
func (n *NoteEvent) GetBase() *string {
	return nil
}
func (n *NoteEvent) GetHead() *string {
	return nil
}
func (n *NoteEvent) GetNumber() *string {
	if n.Attributes != nil && n.Attributes.Number != nil {
		no := strconv.Itoa(*n.Attributes.Number)
		return &no
	}

	if n.PR != nil && n.PR.Number != nil {
		no := strconv.Itoa(*n.PR.Number)
		return &no
	}

	if n.Issue != nil && n.Issue.Number != nil {
		no := strconv.Itoa(*n.Issue.Number)
		return &no
	}

	return nil
}
func (n *NoteEvent) GetAuthor() *string {
	if n.Issue != nil && n.Issue.Author != nil {
		if n.Issue.Author.Login != nil {
			return n.Issue.Author.Login
		}
		return n.Issue.Author.UserName
	}

	if n.PR != nil && n.PR.Author != nil {
		if n.PR.Author.Login != nil {
			return n.PR.Author.Login
		}
		return n.PR.Author.UserName
	}

	return nil
}
func (n *NoteEvent) GetComment() *string {
	if n.Attributes == nil || n.Attributes.Comment == nil {
		return nil
	}
	return n.Attributes.Comment
}
func (n *NoteEvent) GetCommenter() *string {
	if n.User == nil {
		return nil
	}
	if n.User.Login != nil {
		return n.User.Login
	}

	return n.User.UserName
}
func (n *NoteEvent) ListLabels() []*string {
	if len(n.Labels) == 0 {
		return nil
	}

	labels := make([]*string, 0, len(n.Labels))
	for _, p := range n.Labels {
		labels = append(labels, &p.Name)
	}
	return labels
}
