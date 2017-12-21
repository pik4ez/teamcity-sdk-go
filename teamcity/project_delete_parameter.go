package teamcity

import (
	"fmt"
)

func (c *Client) DeleteProjectParameter(projectID, name string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/projects/id:%s/parameters/%s", c.version, projectID, name)
	return c.doRetryRequest("DELETE", path, nil, nil)
}
