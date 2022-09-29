package models

import "time"

const (
	tasksEndpoint = "/tasks"
)

// CreateTaskRequestModel is the model used for creating tasks in todois.
type CreateTaskRequestModel struct {
	ProjectID   string   `json:"project_id,omitempty"`
	SectionID   string   `json:"section_id,omitempty"`
	ParentID    string   `json:"parent_id,omitempty"`
	Content     string   `json:"content"`
	Description string   `json:"description,omitempty"`
	Order       int      `json:"order,omitempty"`
	Labels      []string `json:"labels,omitempty"`
	Priority    int      `json:"priority,omitempty"`
	DueString   string   `json:"due_string,omitempty"`
	DueDate     string   `json:"due_date,omitempty"`
	DueDatetime string   `json:"due_datetime,omitempty"`
	DueLang     string   `json:"due_lang,omitempty"`
	AssigneeID  string   `json:"assignee_id,omitempty"`
}

// Endpoint satisfies the request constraint interface.
func (c CreateTaskRequestModel) Endpoint() string {
	return tasksEndpoint
}

// CreateTaskResponseModel is the model used for unmarshalling created todoist task response.
type CreateTaskResponseModel struct {
	ID           string        `json:"id"`
	ProjectID    string        `json:"project_id"`
	SectionID    interface{}   `json:"section_id"` // unknown type
	ParentID     interface{}   `json:"parent_id"`  // unknown type
	CreatorID    string        `json:"creator_id"`
	AssigneeID   interface{}   `json:"assignee_id"` // unknown type
	AssignerID   interface{}   `json:"assigner_id"` // unknown type
	Content      string        `json:"content"`
	Labels       []interface{} `json:"labels"` // known type but currently not used that's why it's interface{}
	Description  string        `json:"description"`
	Url          string        `json:"url"`
	IsCompleted  bool          `json:"is_completed"`
	Due          Due           `json:"due"`
	CommentCount int           `json:"comment_count"`
	Order        int           `json:"order"`
	Priority     int           `json:"priority"`
	CreatedAt    time.Time     `json:"created_at"`
}

// ResponseModel is no-op it just satisfies the response constraint interface.
func (c CreateTaskResponseModel) ResponseModel() {}

// Due is the model used for unmarshalling todoist task due response.
type Due struct {
	Date        string    `json:"date"`
	String      string    `json:"string"`
	Timezone    string    `json:"timezone"`
	IsRecurring bool      `json:"is_recurring"`
	Datetime    time.Time `json:"datetime"`
}
