package models

const (
	projectsEndpoint = "/projects"
)

// CreateProjectRequestModel is the model used for creating projects in todois.
type CreateProjectRequestModel struct {
	Name       string `json:"name"`
	ParentID   string `json:"parent_id,omitempty"`
	Color      string `json:"color,omitempty"`
	IsFavorite bool   `json:"is_favorite,omitempty"`
	ViewStyle  string `json:"view_style,omitempty"`
}

// Endpoint satisfies the request constraint interface.
func (c CreateProjectRequestModel) Endpoint() string {
	return projectsEndpoint
}

// CreateProjectResponseModel is the model used for unmarshalling created todoist project response.
type CreateProjectResponseModel struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	Url            string      `json:"url"`
	Color          string      `json:"color"`
	ViewStyle      string      `json:"view_style"`
	ParentId       interface{} `json:"parent_id"` // unknown type
	Order          int         `json:"order"`
	CommentCount   int         `json:"comment_count"`
	IsFavorite     bool        `json:"is_favorite"`
	IsInboxProject bool        `json:"is_inbox_project"`
	IsTeamInbox    bool        `json:"is_team_inbox"`
	IsShared       bool        `json:"is_shared"`
}

// ResponseModel is no-op it just satisfies the response constraint interface.
func (c CreateProjectResponseModel) ResponseModel() {}
