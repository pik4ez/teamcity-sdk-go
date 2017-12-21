package teamcity

import (
	"fmt"
)

func (c *Client) DeleteProject(projectID string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/projects/id:%s", c.version, projectID)
	return c.doRetryRequest("DELETE", path, nil, nil)
}
