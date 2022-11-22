package db

import "time"

type GitIssues []struct {
	ID                int         `bson:"id,omitempty"`
	NodeID            string      `bson:"node_id,omitempty"`
	URL               string      `bson:"url,omitempty"`
	RepositoryURL     string      `bson:"repository_url,omitempty"`
	LabelsURL         string      `bson:"labels_url,omitempty"`
	CommentsURL       string      `bson:"comments_url,omitempty"`
	EventsURL         string      `bson:"events_url,omitempty"`
	HTMLURL           string      `bson:"html_url,omitempty"`
	Number            int         `bson:"number,omitempty"`
	State             string      `bson:"state,omitempty"`
	Title             string      `bson:"title,omitempty"`
	Body              string      `bson:"body,omitempty"`
	User              User        `bson:"user,omitempty"`
	Labels            []Labels    `bson:"labels,omitempty"`
	Assignee          Assignee    `bson:"assignee,omitempty"`
	Assignees         []Assignees `bson:"assignees,omitempty"`
	Milestone         Milestone   `bson:"milestone,omitempty"`
	Locked            bool        `bson:"locked,omitempty"`
	ActiveLockReason  string      `bson:"active_lock_reason,omitempty"`
	Comments          int         `bson:"comments,omitempty"`
	PullRequest       PullRequest `bson:"pull_request,omitempty"`
	ClosedAt          interface{} `bson:"closed_at,omitempty"`
	CreatedAt         time.Time   `bson:"created_at,omitempty"`
	UpdatedAt         time.Time   `bson:"updated_at,omitempty"`
	Repository        Repository  `bson:"repository,omitempty"`
	AuthorAssociation string      `bson:"author_association,omitempty"`
}
type User struct {
	Login             string `bson:"login,omitempty"`
	ID                int    `bson:"id,omitempty"`
	NodeID            string `bson:"node_id,omitempty"`
	AvatarURL         string `bson:"avatar_url,omitempty"`
	GravatarID        string `bson:"gravatar_id,omitempty"`
	URL               string `bson:"url,omitempty"`
	HTMLURL           string `bson:"html_url,omitempty"`
	FollowersURL      string `bson:"followers_url,omitempty"`
	FollowingURL      string `bson:"following_url,omitempty"`
	GistsURL          string `bson:"gists_url,omitempty"`
	StarredURL        string `bson:"starred_url,omitempty"`
	SubscriptionsURL  string `bson:"subscriptions_url,omitempty"`
	OrganizationsURL  string `bson:"organizations_url,omitempty"`
	ReposURL          string `bson:"repos_url,omitempty"`
	EventsURL         string `bson:"events_url,omitempty"`
	ReceivedEventsURL string `bson:"received_events_url,omitempty"`
	Type              string `bson:"type,omitempty"`
	SiteAdmin         bool   `bson:"site_admin,omitempty"`
}
type Labels struct {
	ID          int    `bson:"id,omitempty"`
	NodeID      string `bson:"node_id,omitempty"`
	URL         string `bson:"url,omitempty"`
	Name        string `bson:"name,omitempty"`
	Description string `bson:"description,omitempty"`
	Color       string `bson:"color,omitempty"`
	Default     bool   `bson:"default,omitempty"`
}
type Assignee struct {
	Login             string `bson:"login,omitempty"`
	ID                int    `bson:"id,omitempty"`
	NodeID            string `bson:"node_id,omitempty"`
	AvatarURL         string `bson:"avatar_url,omitempty"`
	GravatarID        string `bson:"gravatar_id,omitempty"`
	URL               string `bson:"url,omitempty"`
	HTMLURL           string `bson:"html_url,omitempty"`
	FollowersURL      string `bson:"followers_url,omitempty"`
	FollowingURL      string `bson:"following_url,omitempty"`
	GistsURL          string `bson:"gists_url,omitempty"`
	StarredURL        string `bson:"starred_url,omitempty"`
	SubscriptionsURL  string `bson:"subscriptions_url,omitempty"`
	OrganizationsURL  string `bson:"organizations_url,omitempty"`
	ReposURL          string `bson:"repos_url,omitempty"`
	EventsURL         string `bson:"events_url,omitempty"`
	ReceivedEventsURL string `bson:"received_events_url,omitempty"`
	Type              string `bson:"type,omitempty"`
	SiteAdmin         bool   `bson:"site_admin,omitempty"`
}
type Assignees struct {
	Login             string `bson:"login,omitempty"`
	ID                int    `bson:"id,omitempty"`
	NodeID            string `bson:"node_id,omitempty"`
	AvatarURL         string `bson:"avatar_url,omitempty"`
	GravatarID        string `bson:"gravatar_id,omitempty"`
	URL               string `bson:"url,omitempty"`
	HTMLURL           string `bson:"html_url,omitempty"`
	FollowersURL      string `bson:"followers_url,omitempty"`
	FollowingURL      string `bson:"following_url,omitempty"`
	GistsURL          string `bson:"gists_url,omitempty"`
	StarredURL        string `bson:"starred_url,omitempty"`
	SubscriptionsURL  string `bson:"subscriptions_url,omitempty"`
	OrganizationsURL  string `bson:"organizations_url,omitempty"`
	ReposURL          string `bson:"repos_url,omitempty"`
	EventsURL         string `bson:"events_url,omitempty"`
	ReceivedEventsURL string `bson:"received_events_url,omitempty"`
	Type              string `bson:"type,omitempty"`
	SiteAdmin         bool   `bson:"site_admin,omitempty"`
}
type Creator struct {
	Login             string `bson:"login,omitempty"`
	ID                int    `bson:"id,omitempty"`
	NodeID            string `bson:"node_id,omitempty"`
	AvatarURL         string `bson:"avatar_url,omitempty"`
	GravatarID        string `bson:"gravatar_id,omitempty"`
	URL               string `bson:"url,omitempty"`
	HTMLURL           string `bson:"html_url,omitempty"`
	FollowersURL      string `bson:"followers_url,omitempty"`
	FollowingURL      string `bson:"following_url,omitempty"`
	GistsURL          string `bson:"gists_url,omitempty"`
	StarredURL        string `bson:"starred_url,omitempty"`
	SubscriptionsURL  string `bson:"subscriptions_url,omitempty"`
	OrganizationsURL  string `bson:"organizations_url,omitempty"`
	ReposURL          string `bson:"repos_url,omitempty"`
	EventsURL         string `bson:"events_url,omitempty"`
	ReceivedEventsURL string `bson:"received_events_url,omitempty"`
	Type              string `bson:"type,omitempty"`
	SiteAdmin         bool   `bson:"site_admin,omitempty"`
}
type Milestone struct {
	URL          string    `bson:"url,omitempty"`
	HTMLURL      string    `bson:"html_url,omitempty"`
	LabelsURL    string    `bson:"labels_url,omitempty"`
	ID           int       `bson:"id,omitempty"`
	NodeID       string    `bson:"node_id,omitempty"`
	Number       int       `bson:"number,omitempty"`
	State        string    `bson:"state,omitempty"`
	Title        string    `bson:"title,omitempty"`
	Description  string    `bson:"description,omitempty"`
	Creator      Creator   `bson:"creator,omitempty"`
	OpenIssues   int       `bson:"open_issues,omitempty"`
	ClosedIssues int       `bson:"closed_issues,omitempty"`
	CreatedAt    time.Time `bson:"created_at,omitempty"`
	UpdatedAt    time.Time `bson:"updated_at,omitempty"`
	ClosedAt     time.Time `bson:"closed_at,omitempty"`
	DueOn        time.Time `bson:"due_on,omitempty"`
}
type PullRequest struct {
	URL      string `bson:"url,omitempty"`
	HTMLURL  string `bson:"html_url,omitempty"`
	DiffURL  string `bson:"diff_url,omitempty"`
	PatchURL string `bson:"patch_url,omitempty"`
}
type Owner struct {
	Login             string `bson:"login,omitempty"`
	ID                int    `bson:"id,omitempty"`
	NodeID            string `bson:"node_id,omitempty"`
	AvatarURL         string `bson:"avatar_url,omitempty"`
	GravatarID        string `bson:"gravatar_id,omitempty"`
	URL               string `bson:"url,omitempty"`
	HTMLURL           string `bson:"html_url,omitempty"`
	FollowersURL      string `bson:"followers_url,omitempty"`
	FollowingURL      string `bson:"following_url,omitempty"`
	GistsURL          string `bson:"gists_url,omitempty"`
	StarredURL        string `bson:"starred_url,omitempty"`
	SubscriptionsURL  string `bson:"subscriptions_url,omitempty"`
	OrganizationsURL  string `bson:"organizations_url,omitempty"`
	ReposURL          string `bson:"repos_url,omitempty"`
	EventsURL         string `bson:"events_url,omitempty"`
	ReceivedEventsURL string `bson:"received_events_url,omitempty"`
	Type              string `bson:"type,omitempty"`
	SiteAdmin         bool   `bson:"site_admin,omitempty"`
}
type Permissions struct {
	Admin bool `bson:"admin,omitempty"`
	Push  bool `bson:"push,omitempty"`
	Pull  bool `bson:"pull,omitempty"`
}
type License struct {
	Key     string `bson:"key,omitempty"`
	Name    string `bson:"name,omitempty"`
	URL     string `bson:"url,omitempty"`
	SpdxID  string `bson:"spdx_id,omitempty"`
	NodeID  string `bson:"node_id,omitempty"`
	HTMLURL string `bson:"html_url,omitempty"`
}
type Repository struct {
	ID                  int         `bson:"id,omitempty"`
	NodeID              string      `bson:"node_id,omitempty"`
	Name                string      `bson:"name,omitempty"`
	FullName            string      `bson:"full_name,omitempty"`
	Owner               Owner       `bson:"owner,omitempty"`
	Private             bool        `bson:"private,omitempty"`
	HTMLURL             string      `bson:"html_url,omitempty"`
	Description         string      `bson:"description,omitempty"`
	Fork                bool        `bson:"fork,omitempty"`
	URL                 string      `bson:"url,omitempty"`
	ArchiveURL          string      `bson:"archive_url,omitempty"`
	AssigneesURL        string      `bson:"assignees_url,omitempty"`
	BlobsURL            string      `bson:"blobs_url,omitempty"`
	BranchesURL         string      `bson:"branches_url,omitempty"`
	CollaboratorsURL    string      `bson:"collaborators_url,omitempty"`
	CommentsURL         string      `bson:"comments_url,omitempty"`
	CommitsURL          string      `bson:"commits_url,omitempty"`
	CompareURL          string      `bson:"compare_url,omitempty"`
	ContentsURL         string      `bson:"contents_url,omitempty"`
	ContributorsURL     string      `bson:"contributors_url,omitempty"`
	DeploymentsURL      string      `bson:"deployments_url,omitempty"`
	DownloadsURL        string      `bson:"downloads_url,omitempty"`
	EventsURL           string      `bson:"events_url,omitempty"`
	ForksURL            string      `bson:"forks_url,omitempty"`
	GitCommitsURL       string      `bson:"git_commits_url,omitempty"`
	GitRefsURL          string      `bson:"git_refs_url,omitempty"`
	GitTagsURL          string      `bson:"git_tags_url,omitempty"`
	GitURL              string      `bson:"git_url,omitempty"`
	IssueCommentURL     string      `bson:"issue_comment_url,omitempty"`
	IssueEventsURL      string      `bson:"issue_events_url,omitempty"`
	IssuesURL           string      `bson:"issues_url,omitempty"`
	KeysURL             string      `bson:"keys_url,omitempty"`
	LabelsURL           string      `bson:"labels_url,omitempty"`
	LanguagesURL        string      `bson:"languages_url,omitempty"`
	MergesURL           string      `bson:"merges_url,omitempty"`
	MilestonesURL       string      `bson:"milestones_url,omitempty"`
	NotificationsURL    string      `bson:"notifications_url,omitempty"`
	PullsURL            string      `bson:"pulls_url,omitempty"`
	ReleasesURL         string      `bson:"releases_url,omitempty"`
	SSHURL              string      `bson:"ssh_url,omitempty"`
	StargazersURL       string      `bson:"stargazers_url,omitempty"`
	StatusesURL         string      `bson:"statuses_url,omitempty"`
	SubscribersURL      string      `bson:"subscribers_url,omitempty"`
	SubscriptionURL     string      `bson:"subscription_url,omitempty"`
	TagsURL             string      `bson:"tags_url,omitempty"`
	TeamsURL            string      `bson:"teams_url,omitempty"`
	TreesURL            string      `bson:"trees_url,omitempty"`
	CloneURL            string      `bson:"clone_url,omitempty"`
	MirrorURL           string      `bson:"mirror_url,omitempty"`
	HooksURL            string      `bson:"hooks_url,omitempty"`
	SvnURL              string      `bson:"svn_url,omitempty"`
	Homepage            string      `bson:"homepage,omitempty"`
	Language            interface{} `bson:"language,omitempty"`
	ForksCount          int         `bson:"forks_count,omitempty"`
	StargazersCount     int         `bson:"stargazers_count,omitempty"`
	WatchersCount       int         `bson:"watchers_count,omitempty"`
	Size                int         `bson:"size,omitempty"`
	DefaultBranch       string      `bson:"default_branch,omitempty"`
	OpenIssuesCount     int         `bson:"open_issues_count,omitempty"`
	IsTemplate          bool        `bson:"is_template,omitempty"`
	Topics              []string    `bson:"topics,omitempty"`
	HasIssues           bool        `bson:"has_issues,omitempty"`
	HasProjects         bool        `bson:"has_projects,omitempty"`
	HasWiki             bool        `bson:"has_wiki,omitempty"`
	HasPages            bool        `bson:"has_pages,omitempty"`
	HasDownloads        bool        `bson:"has_downloads,omitempty"`
	Archived            bool        `bson:"archived,omitempty"`
	Disabled            bool        `bson:"disabled,omitempty"`
	Visibility          string      `bson:"visibility,omitempty"`
	PushedAt            time.Time   `bson:"pushed_at,omitempty"`
	CreatedAt           time.Time   `bson:"created_at,omitempty"`
	UpdatedAt           time.Time   `bson:"updated_at,omitempty"`
	Permissions         Permissions `bson:"permissions,omitempty"`
	AllowRebaseMerge    bool        `bson:"allow_rebase_merge,omitempty"`
	TemplateRepository  interface{} `bson:"template_repository,omitempty"`
	TempCloneToken      string      `bson:"temp_clone_token,omitempty"`
	AllowSquashMerge    bool        `bson:"allow_squash_merge,omitempty"`
	AllowAutoMerge      bool        `bson:"allow_auto_merge,omitempty"`
	DeleteBranchOnMerge bool        `bson:"delete_branch_on_merge,omitempty"`
	AllowMergeCommit    bool        `bson:"allow_merge_commit,omitempty"`
	SubscribersCount    int         `bson:"subscribers_count,omitempty"`
	NetworkCount        int         `bson:"network_count,omitempty"`
	License             License     `bson:"license,omitempty"`
	Forks               int         `bson:"forks,omitempty"`
	OpenIssues          int         `bson:"open_issues,omitempty"`
	Watchers            int         `bson:"watchers,omitempty"`
}
