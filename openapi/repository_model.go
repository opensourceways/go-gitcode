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

// Repository represents a GitHub repository.
type Repository struct {
	ID                        *int64                 `json:"id,omitempty"`
	NodeID                    *string                `json:"node_id,omitempty"`
	Owner                     *User                  `json:"owner,omitempty"`
	Name                      *string                `json:"name,omitempty"`
	Path                      *string                `json:"path,omitempty"`
	FullName                  *string                `json:"full_name,omitempty"`
	Description               *string                `json:"description,omitempty"`
	Homepage                  *string                `json:"homepage,omitempty"`
	DefaultBranch             *string                `json:"default_branch,omitempty"`
	MasterBranch              *string                `json:"master_branch,omitempty"`
	CreatedAt                 *Timestamp             `json:"created_at,omitempty"`
	PushedAt                  *Timestamp             `json:"pushed_at,omitempty"`
	UpdatedAt                 *Timestamp             `json:"updated_at,omitempty"`
	HTMLURL                   *string                `json:"html_url,omitempty"`
	CloneURL                  *string                `json:"clone_url,omitempty"`
	GitURL                    *string                `json:"git_url,omitempty"`
	MirrorURL                 *string                `json:"mirror_url,omitempty"`
	SSHURL                    *string                `json:"ssh_url,omitempty"`
	SVNURL                    *string                `json:"svn_url,omitempty"`
	Language                  *string                `json:"language,omitempty"`
	Fork                      *bool                  `json:"fork,omitempty"`
	ForksCount                *int                   `json:"forks_count,omitempty"`
	NetworkCount              *int                   `json:"network_count,omitempty"`
	OpenIssuesCount           *int                   `json:"open_issues_count,omitempty"`
	OpenIssues                *int                   `json:"open_issues,omitempty"` // Deprecated: Replaced by OpenIssuesCount. For backward compatibility OpenIssues is still populated.
	StargazersCount           *int                   `json:"stargazers_count,omitempty"`
	SubscribersCount          *int                   `json:"subscribers_count,omitempty"`
	WatchersCount             *int                   `json:"watchers_count,omitempty"` // Deprecated: Replaced by StargazersCount. For backward compatibility WatchersCount is still populated.
	Watchers                  *int                   `json:"watchers,omitempty"`       // Deprecated: Replaced by StargazersCount. For backward compatibility Watchers is still populated.
	Size                      *int                   `json:"size,omitempty"`
	AutoInit                  *bool                  `json:"auto_init,omitempty"`
	Parent                    *Repository            `json:"parent,omitempty"`
	Source                    *Repository            `json:"source,omitempty"`
	TemplateRepository        *Repository            `json:"template_repository,omitempty"`
	Organization              *Organization          `json:"organization,omitempty"`
	Permissions               map[string]bool        `json:"permissions,omitempty"`
	AllowRebaseMerge          *bool                  `json:"allow_rebase_merge,omitempty"`
	AllowUpdateBranch         *bool                  `json:"allow_update_branch,omitempty"`
	AllowSquashMerge          *bool                  `json:"allow_squash_merge,omitempty"`
	AllowMergeCommit          *bool                  `json:"allow_merge_commit,omitempty"`
	AllowAutoMerge            *bool                  `json:"allow_auto_merge,omitempty"`
	AllowForking              *bool                  `json:"allow_forking,omitempty"`
	WebCommitSignoffRequired  *bool                  `json:"web_commit_signoff_required,omitempty"`
	DeleteBranchOnMerge       *bool                  `json:"delete_branch_on_merge,omitempty"`
	UseSquashPRTitleAsDefault *bool                  `json:"use_squash_pr_title_as_default,omitempty"`
	SquashMergeCommitTitle    *string                `json:"squash_merge_commit_title,omitempty"`   // Can be one of: "PR_TITLE", "COMMIT_OR_PR_TITLE"
	SquashMergeCommitMessage  *string                `json:"squash_merge_commit_message,omitempty"` // Can be one of: "PR_BODY", "COMMIT_MESSAGES", "BLANK"
	MergeCommitTitle          *string                `json:"merge_commit_title,omitempty"`          // Can be one of: "PR_TITLE", "MERGE_MESSAGE"
	MergeCommitMessage        *string                `json:"merge_commit_message,omitempty"`        // Can be one of: "PR_BODY", "PR_TITLE", "BLANK"
	Topics                    []string               `json:"topics,omitempty"`
	CustomProperties          map[string]interface{} `json:"custom_properties,omitempty"`
	Archived                  *bool                  `json:"archived,omitempty"`
	Disabled                  *bool                  `json:"disabled,omitempty"`

	// Additional mutable fields when creating and editing a repository
	Private           *bool   `json:"private,omitempty"`
	HasIssues         *bool   `json:"has_issues,omitempty"`
	HasWiki           *bool   `json:"has_wiki,omitempty"`
	HasPages          *bool   `json:"has_pages,omitempty"`
	HasProjects       *bool   `json:"has_projects,omitempty"`
	HasDownloads      *bool   `json:"has_downloads,omitempty"`
	HasDiscussions    *bool   `json:"has_discussions,omitempty"`
	IsTemplate        *bool   `json:"is_template,omitempty"`
	LicenseTemplate   *string `json:"license_template,omitempty"`
	GitignoreTemplate *string `json:"gitignore_template,omitempty"`

	// Creating an organization repository. Required for non-owners.
	TeamID *int64 `json:"team_id,omitempty"`

	// api URLs
	URL              *string `json:"url,omitempty"`
	ArchiveURL       *string `json:"archive_url,omitempty"`
	AssigneesURL     *string `json:"assignees_url,omitempty"`
	BlobsURL         *string `json:"blobs_url,omitempty"`
	BranchesURL      *string `json:"branches_url,omitempty"`
	CollaboratorsURL *string `json:"collaborators_url,omitempty"`
	CommentsURL      *string `json:"comments_url,omitempty"`
	CommitsURL       *string `json:"commits_url,omitempty"`
	CompareURL       *string `json:"compare_url,omitempty"`
	ContentsURL      *string `json:"contents_url,omitempty"`
	ContributorsURL  *string `json:"contributors_url,omitempty"`
	DeploymentsURL   *string `json:"deployments_url,omitempty"`
	DownloadsURL     *string `json:"downloads_url,omitempty"`
	EventsURL        *string `json:"events_url,omitempty"`
	ForksURL         *string `json:"forks_url,omitempty"`
	GitCommitsURL    *string `json:"git_commits_url,omitempty"`
	GitRefsURL       *string `json:"git_refs_url,omitempty"`
	GitTagsURL       *string `json:"git_tags_url,omitempty"`
	HooksURL         *string `json:"hooks_url,omitempty"`
	IssueCommentURL  *string `json:"issue_comment_url,omitempty"`
	IssueEventsURL   *string `json:"issue_events_url,omitempty"`
	IssuesURL        *string `json:"issues_url,omitempty"`
	KeysURL          *string `json:"keys_url,omitempty"`
	LabelsURL        *string `json:"labels_url,omitempty"`
	LanguagesURL     *string `json:"languages_url,omitempty"`
	MergesURL        *string `json:"merges_url,omitempty"`
	MilestonesURL    *string `json:"milestones_url,omitempty"`
	NotificationsURL *string `json:"notifications_url,omitempty"`
	PullsURL         *string `json:"pulls_url,omitempty"`
	ReleasesURL      *string `json:"releases_url,omitempty"`
	StargazersURL    *string `json:"stargazers_url,omitempty"`
	StatusesURL      *string `json:"statuses_url,omitempty"`
	SubscribersURL   *string `json:"subscribers_url,omitempty"`
	SubscriptionURL  *string `json:"subscription_url,omitempty"`
	TagsURL          *string `json:"tags_url,omitempty"`
	TreesURL         *string `json:"trees_url,omitempty"`
	TeamsURL         *string `json:"teams_url,omitempty"`

	// Visibility is only used for Create and Edit endpoints. The visibility field
	// overrides the field parameter when both are used.
	// Can be one of public, private or internal.
	Visibility *string `json:"visibility,omitempty"`

	// RoleName is only returned by the api 'check team permissions for a repository'.
	// See: teams.go (IsTeamRepoByID) https://docs.github.com/rest/teams/teams#check-team-permissions-for-a-repository
	RoleName *string `json:"role_name,omitempty"`
}

type RepositoryCommit struct {
	SHA         *string     `json:"sha,omitempty"`
	Commit      *Commit     `json:"commit,omitempty"`
	Author      *CommitUser `json:"author,omitempty"`
	Committer   *CommitUser `json:"committer,omitempty"`
	Parents     *Commit     `json:"parents,omitempty"`
	HTMLURL     *string     `json:"html_url,omitempty"`
	URL         *string     `json:"url,omitempty"`
	CommentsURL *string     `json:"comments_url,omitempty"`
}

type Commit struct {
	SHA         *string     `json:"sha,omitempty"`
	Author      *CommitUser `json:"author,omitempty"`
	Committer   *CommitUser `json:"committer,omitempty"`
	Message     *string     `json:"message,omitempty"`
	Parents     *Commit     `json:"parents,omitempty"`
	HTMLURL     *string     `json:"html_url,omitempty"`
	URL         *string     `json:"url,omitempty"`
	CommentsURL *int        `json:"comments_url,omitempty"`
}

type CommitUser struct {
	Login *string    `json:"login,omitempty"`
	Name  *string    `json:"name,omitempty"`
	Email *string    `json:"email,omitempty"`
	Date  *Timestamp `json:"date,omitempty"`
}

type CommitPatch struct {
	Diff        *string `json:"diff,omitempty"`
	OldPath     *string `json:"old_path,omitempty"`
	NewPath     *string `json:"new_path,omitempty"`
	NewFile     *bool   `json:"new_file,omitempty"`
	RenamedFile *bool   `json:"renamed_file,omitempty"`
	DeletedFile *bool   `json:"deleted_file,omitempty"`
	TooLarge    *bool   `json:"too_large,omitempty"`
}

type CommitFile struct {
	SHA              *string      `json:"sha,omitempty"`
	Filename         *string      `json:"filename,omitempty"`
	Additions        *int         `json:"additions,omitempty"`
	Deletions        *int         `json:"deletions,omitempty"`
	Changes          *int         `json:"changes,omitempty"`
	Status           *string      `json:"status,omitempty"`
	Patch            *CommitPatch `json:"patch,omitempty"`
	BlobURL          *string      `json:"blob_url,omitempty"`
	RawURL           *string      `json:"raw_url,omitempty"`
	ContentsURL      *string      `json:"contents_url,omitempty"`
	PreviousFilename *string      `json:"previous_filename,omitempty"`
}

// RepositoryContent represents a file or directory in a github repository.
type RepositoryContent struct {
	Type *string `json:"type,omitempty"`
	// Target is only set if the type is "symlink" and the target is not a normal file.
	// If Target is set, Path will be the symlink path.
	Target   *string `json:"target,omitempty"`
	Encoding *string `json:"encoding,omitempty"`
	Size     *int64  `json:"size,omitempty"`
	Name     *string `json:"name,omitempty"`
	Path     *string `json:"path,omitempty"`
	// Content contains the actual file content, which may be encoded.
	// Callers should call GetContent which will decode the content if
	// necessary.
	Content         *string `json:"content,omitempty"`
	SHA             *string `json:"sha,omitempty"`
	URL             *string `json:"url,omitempty"`
	GitURL          *string `json:"git_url,omitempty"`
	HTMLURL         *string `json:"html_url,omitempty"`
	DownloadURL     *string `json:"download_url,omitempty"`
	SubmoduleGitURL *string `json:"submodule_git_url,omitempty"`
}

// Contributor represents a repository contributor
type Contributor struct {
	Contributions *int64  `json:"contributions,omitempty"`
	Name          *string `json:"name,omitempty"`
	Email         *string `json:"email,omitempty"`
}

// Branch represents a repository branch
type Branch struct {
	Name      *string           `json:"name,omitempty"`
	Commit    *RepositoryCommit `json:"commit,omitempty"`
	Protected *bool             `json:"protected,omitempty"`
}
