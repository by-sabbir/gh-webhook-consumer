package hooks

import "time"

type GitIssues []struct {
	ID                int         `json:"id,omitempty"`
	NodeID            string      `json:"node_id,omitempty"`
	URL               string      `json:"url,omitempty"`
	RepositoryURL     string      `json:"repository_url,omitempty"`
	LabelsURL         string      `json:"labels_url,omitempty"`
	CommentsURL       string      `json:"comments_url,omitempty"`
	EventsURL         string      `json:"events_url,omitempty"`
	HTMLURL           string      `json:"html_url,omitempty"`
	Number            int         `json:"number,omitempty"`
	State             string      `json:"state,omitempty"`
	Title             string      `json:"title,omitempty"`
	Body              string      `json:"body,omitempty"`
	User              User        `json:"user,omitempty"`
	Labels            []Labels    `json:"labels,omitempty"`
	Assignee          Assignee    `json:"assignee,omitempty"`
	Assignees         []Assignees `json:"assignees,omitempty"`
	Milestone         Milestone   `json:"milestone,omitempty"`
	Locked            bool        `json:"locked,omitempty"`
	ActiveLockReason  string      `json:"active_lock_reason,omitempty"`
	Comments          int         `json:"comments,omitempty"`
	PullRequest       PullRequest `json:"pull_request,omitempty"`
	ClosedAt          interface{} `json:"closed_at,omitempty"`
	CreatedAt         time.Time   `json:"created_at,omitempty"`
	UpdatedAt         time.Time   `json:"updated_at,omitempty"`
	Repository        Repository  `json:"repository,omitempty"`
	AuthorAssociation string      `json:"author_association,omitempty"`
}
type User struct {
	Login             string `json:"login,omitempty"`
	ID                int    `json:"id,omitempty"`
	NodeID            string `json:"node_id,omitempty"`
	AvatarURL         string `json:"avatar_url,omitempty"`
	GravatarID        string `json:"gravatar_id,omitempty"`
	URL               string `json:"url,omitempty"`
	HTMLURL           string `json:"html_url,omitempty"`
	FollowersURL      string `json:"followers_url,omitempty"`
	FollowingURL      string `json:"following_url,omitempty"`
	GistsURL          string `json:"gists_url,omitempty"`
	StarredURL        string `json:"starred_url,omitempty"`
	SubscriptionsURL  string `json:"subscriptions_url,omitempty"`
	OrganizationsURL  string `json:"organizations_url,omitempty"`
	ReposURL          string `json:"repos_url,omitempty"`
	EventsURL         string `json:"events_url,omitempty"`
	ReceivedEventsURL string `json:"received_events_url,omitempty"`
	Type              string `json:"type,omitempty"`
	SiteAdmin         bool   `json:"site_admin,omitempty"`
}
type Labels struct {
	ID          int    `json:"id,omitempty"`
	NodeID      string `json:"node_id,omitempty"`
	URL         string `json:"url,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Color       string `json:"color,omitempty"`
	Default     bool   `json:"default,omitempty"`
}
type Assignee struct {
	Login             string `json:"login,omitempty"`
	ID                int    `json:"id,omitempty"`
	NodeID            string `json:"node_id,omitempty"`
	AvatarURL         string `json:"avatar_url,omitempty"`
	GravatarID        string `json:"gravatar_id,omitempty"`
	URL               string `json:"url,omitempty"`
	HTMLURL           string `json:"html_url,omitempty"`
	FollowersURL      string `json:"followers_url,omitempty"`
	FollowingURL      string `json:"following_url,omitempty"`
	GistsURL          string `json:"gists_url,omitempty"`
	StarredURL        string `json:"starred_url,omitempty"`
	SubscriptionsURL  string `json:"subscriptions_url,omitempty"`
	OrganizationsURL  string `json:"organizations_url,omitempty"`
	ReposURL          string `json:"repos_url,omitempty"`
	EventsURL         string `json:"events_url,omitempty"`
	ReceivedEventsURL string `json:"received_events_url,omitempty"`
	Type              string `json:"type,omitempty"`
	SiteAdmin         bool   `json:"site_admin,omitempty"`
}
type Assignees struct {
	Login             string `json:"login,omitempty"`
	ID                int    `json:"id,omitempty"`
	NodeID            string `json:"node_id,omitempty"`
	AvatarURL         string `json:"avatar_url,omitempty"`
	GravatarID        string `json:"gravatar_id,omitempty"`
	URL               string `json:"url,omitempty"`
	HTMLURL           string `json:"html_url,omitempty"`
	FollowersURL      string `json:"followers_url,omitempty"`
	FollowingURL      string `json:"following_url,omitempty"`
	GistsURL          string `json:"gists_url,omitempty"`
	StarredURL        string `json:"starred_url,omitempty"`
	SubscriptionsURL  string `json:"subscriptions_url,omitempty"`
	OrganizationsURL  string `json:"organizations_url,omitempty"`
	ReposURL          string `json:"repos_url,omitempty"`
	EventsURL         string `json:"events_url,omitempty"`
	ReceivedEventsURL string `json:"received_events_url,omitempty"`
	Type              string `json:"type,omitempty"`
	SiteAdmin         bool   `json:"site_admin,omitempty"`
}
type Creator struct {
	Login             string `json:"login,omitempty"`
	ID                int    `json:"id,omitempty"`
	NodeID            string `json:"node_id,omitempty"`
	AvatarURL         string `json:"avatar_url,omitempty"`
	GravatarID        string `json:"gravatar_id,omitempty"`
	URL               string `json:"url,omitempty"`
	HTMLURL           string `json:"html_url,omitempty"`
	FollowersURL      string `json:"followers_url,omitempty"`
	FollowingURL      string `json:"following_url,omitempty"`
	GistsURL          string `json:"gists_url,omitempty"`
	StarredURL        string `json:"starred_url,omitempty"`
	SubscriptionsURL  string `json:"subscriptions_url,omitempty"`
	OrganizationsURL  string `json:"organizations_url,omitempty"`
	ReposURL          string `json:"repos_url,omitempty"`
	EventsURL         string `json:"events_url,omitempty"`
	ReceivedEventsURL string `json:"received_events_url,omitempty"`
	Type              string `json:"type,omitempty"`
	SiteAdmin         bool   `json:"site_admin,omitempty"`
}
type Milestone struct {
	URL          string    `json:"url,omitempty"`
	HTMLURL      string    `json:"html_url,omitempty"`
	LabelsURL    string    `json:"labels_url,omitempty"`
	ID           int       `json:"id,omitempty"`
	NodeID       string    `json:"node_id,omitempty"`
	Number       int       `json:"number,omitempty"`
	State        string    `json:"state,omitempty"`
	Title        string    `json:"title,omitempty"`
	Description  string    `json:"description,omitempty"`
	Creator      Creator   `json:"creator,omitempty"`
	OpenIssues   int       `json:"open_issues,omitempty"`
	ClosedIssues int       `json:"closed_issues,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	ClosedAt     time.Time `json:"closed_at,omitempty"`
	DueOn        time.Time `json:"due_on,omitempty"`
}
type PullRequest struct {
	URL      string `json:"url,omitempty"`
	HTMLURL  string `json:"html_url,omitempty"`
	DiffURL  string `json:"diff_url,omitempty"`
	PatchURL string `json:"patch_url,omitempty"`
}
type Owner struct {
	Login             string `json:"login,omitempty"`
	ID                int    `json:"id,omitempty"`
	NodeID            string `json:"node_id,omitempty"`
	AvatarURL         string `json:"avatar_url,omitempty"`
	GravatarID        string `json:"gravatar_id,omitempty"`
	URL               string `json:"url,omitempty"`
	HTMLURL           string `json:"html_url,omitempty"`
	FollowersURL      string `json:"followers_url,omitempty"`
	FollowingURL      string `json:"following_url,omitempty"`
	GistsURL          string `json:"gists_url,omitempty"`
	StarredURL        string `json:"starred_url,omitempty"`
	SubscriptionsURL  string `json:"subscriptions_url,omitempty"`
	OrganizationsURL  string `json:"organizations_url,omitempty"`
	ReposURL          string `json:"repos_url,omitempty"`
	EventsURL         string `json:"events_url,omitempty"`
	ReceivedEventsURL string `json:"received_events_url,omitempty"`
	Type              string `json:"type,omitempty"`
	SiteAdmin         bool   `json:"site_admin,omitempty"`
}
type Permissions struct {
	Admin bool `json:"admin,omitempty"`
	Push  bool `json:"push,omitempty"`
	Pull  bool `json:"pull,omitempty"`
}
type License struct {
	Key     string `json:"key,omitempty"`
	Name    string `json:"name,omitempty"`
	URL     string `json:"url,omitempty"`
	SpdxID  string `json:"spdx_id,omitempty"`
	NodeID  string `json:"node_id,omitempty"`
	HTMLURL string `json:"html_url,omitempty"`
}
type Repository struct {
	ID                  int         `json:"id,omitempty"`
	NodeID              string      `json:"node_id,omitempty"`
	Name                string      `json:"name,omitempty"`
	FullName            string      `json:"full_name,omitempty"`
	Owner               Owner       `json:"owner,omitempty"`
	Private             bool        `json:"private,omitempty"`
	HTMLURL             string      `json:"html_url,omitempty"`
	Description         string      `json:"description,omitempty"`
	Fork                bool        `json:"fork,omitempty"`
	URL                 string      `json:"url,omitempty"`
	ArchiveURL          string      `json:"archive_url,omitempty"`
	AssigneesURL        string      `json:"assignees_url,omitempty"`
	BlobsURL            string      `json:"blobs_url,omitempty"`
	BranchesURL         string      `json:"branches_url,omitempty"`
	CollaboratorsURL    string      `json:"collaborators_url,omitempty"`
	CommentsURL         string      `json:"comments_url,omitempty"`
	CommitsURL          string      `json:"commits_url,omitempty"`
	CompareURL          string      `json:"compare_url,omitempty"`
	ContentsURL         string      `json:"contents_url,omitempty"`
	ContributorsURL     string      `json:"contributors_url,omitempty"`
	DeploymentsURL      string      `json:"deployments_url,omitempty"`
	DownloadsURL        string      `json:"downloads_url,omitempty"`
	EventsURL           string      `json:"events_url,omitempty"`
	ForksURL            string      `json:"forks_url,omitempty"`
	GitCommitsURL       string      `json:"git_commits_url,omitempty"`
	GitRefsURL          string      `json:"git_refs_url,omitempty"`
	GitTagsURL          string      `json:"git_tags_url,omitempty"`
	GitURL              string      `json:"git_url,omitempty"`
	IssueCommentURL     string      `json:"issue_comment_url,omitempty"`
	IssueEventsURL      string      `json:"issue_events_url,omitempty"`
	IssuesURL           string      `json:"issues_url,omitempty"`
	KeysURL             string      `json:"keys_url,omitempty"`
	LabelsURL           string      `json:"labels_url,omitempty"`
	LanguagesURL        string      `json:"languages_url,omitempty"`
	MergesURL           string      `json:"merges_url,omitempty"`
	MilestonesURL       string      `json:"milestones_url,omitempty"`
	NotificationsURL    string      `json:"notifications_url,omitempty"`
	PullsURL            string      `json:"pulls_url,omitempty"`
	ReleasesURL         string      `json:"releases_url,omitempty"`
	SSHURL              string      `json:"ssh_url,omitempty"`
	StargazersURL       string      `json:"stargazers_url,omitempty"`
	StatusesURL         string      `json:"statuses_url,omitempty"`
	SubscribersURL      string      `json:"subscribers_url,omitempty"`
	SubscriptionURL     string      `json:"subscription_url,omitempty"`
	TagsURL             string      `json:"tags_url,omitempty"`
	TeamsURL            string      `json:"teams_url,omitempty"`
	TreesURL            string      `json:"trees_url,omitempty"`
	CloneURL            string      `json:"clone_url,omitempty"`
	MirrorURL           string      `json:"mirror_url,omitempty"`
	HooksURL            string      `json:"hooks_url,omitempty"`
	SvnURL              string      `json:"svn_url,omitempty"`
	Homepage            string      `json:"homepage,omitempty"`
	Language            interface{} `json:"language,omitempty"`
	ForksCount          int         `json:"forks_count,omitempty"`
	StargazersCount     int         `json:"stargazers_count,omitempty"`
	WatchersCount       int         `json:"watchers_count,omitempty"`
	Size                int         `json:"size,omitempty"`
	DefaultBranch       string      `json:"default_branch,omitempty"`
	OpenIssuesCount     int         `json:"open_issues_count,omitempty"`
	IsTemplate          bool        `json:"is_template,omitempty"`
	Topics              []string    `json:"topics,omitempty"`
	HasIssues           bool        `json:"has_issues,omitempty"`
	HasProjects         bool        `json:"has_projects,omitempty"`
	HasWiki             bool        `json:"has_wiki,omitempty"`
	HasPages            bool        `json:"has_pages,omitempty"`
	HasDownloads        bool        `json:"has_downloads,omitempty"`
	Archived            bool        `json:"archived,omitempty"`
	Disabled            bool        `json:"disabled,omitempty"`
	Visibility          string      `json:"visibility,omitempty"`
	PushedAt            time.Time   `json:"pushed_at,omitempty"`
	CreatedAt           time.Time   `json:"created_at,omitempty"`
	UpdatedAt           time.Time   `json:"updated_at,omitempty"`
	Permissions         Permissions `json:"permissions,omitempty"`
	AllowRebaseMerge    bool        `json:"allow_rebase_merge,omitempty"`
	TemplateRepository  interface{} `json:"template_repository,omitempty"`
	TempCloneToken      string      `json:"temp_clone_token,omitempty"`
	AllowSquashMerge    bool        `json:"allow_squash_merge,omitempty"`
	AllowAutoMerge      bool        `json:"allow_auto_merge,omitempty"`
	DeleteBranchOnMerge bool        `json:"delete_branch_on_merge,omitempty"`
	AllowMergeCommit    bool        `json:"allow_merge_commit,omitempty"`
	SubscribersCount    int         `json:"subscribers_count,omitempty"`
	NetworkCount        int         `json:"network_count,omitempty"`
	License             License     `json:"license,omitempty"`
	Forks               int         `json:"forks,omitempty"`
	OpenIssues          int         `json:"open_issues,omitempty"`
	Watchers            int         `json:"watchers,omitempty"`
}
