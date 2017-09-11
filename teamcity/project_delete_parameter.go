package teamcity

import (
	"fmt"
)

func (c *Client) DeleteProjectParameter(projectID, name string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/projects/id:%s/parameters/%s", projectID, name)
	return c.doRetryRequest("DELETE", path, nil, nil)
}
