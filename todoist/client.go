package todoist

import (
	"fmt"
	"io"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/petrovskiborislav/cyberpunk2077-todoist-quests/todoist/models"
)

const (
	baseURL              = "https://api.todoist.com/rest/v2"
	token                = "put_your_token_here"
	defaultRetryWaitTime = 5 * time.Second
)

// Client is the interface that wraps the basic methods of the Todoist API.
type Client interface {
	Create(request models.TodoistRequestModelConstraint, result models.TodoistResponseModelConstraint) error
}

// NewClient returns a new instance of the Todoist API client.
func NewClient() Client {
	httReq := resty.New().
		SetBaseURL(baseURL).
		SetAuthToken(token).
		AddRetryAfterErrorCondition().
		SetRetryWaitTime(defaultRetryWaitTime).
		SetRetryCount(3).R()
	return &client{httpRequest: httReq}
}

type client struct {
	httpRequest *resty.Request
}

func (c *client) Create(request models.TodoistRequestModelConstraint, result models.TodoistResponseModelConstraint) error {
	res, err := c.httpRequest.SetBody(request).SetResult(result).Post(request.Endpoint())
	if err != nil {
		return err
	}

	if res.IsError() {
		if res.Error() == nil {
			resp, err := io.ReadAll(res.RawBody())
			if err != nil {
				return err
			}
			return fmt.Errorf("error on creating todoist task: %v - status code: %d", string(resp), res.StatusCode())
		}
		return fmt.Errorf("error on creating todoist task: %v", res.Error())
	}

	return nil
}
