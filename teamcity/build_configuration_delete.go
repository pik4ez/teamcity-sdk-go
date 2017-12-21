package teamcity

import (
	"fmt"
)

func (c *Client) DeleteBuildConfiguration(buildConfID string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/buildTypes/id:%s", c.version, buildConfID)
	return c.doRetryRequest("DELETE", path, nil, nil)
}
