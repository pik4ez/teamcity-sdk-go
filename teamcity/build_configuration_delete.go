package teamcity

import (
	"fmt"
)

func (c *Client) DeleteBuildConfiguration(buildConfID string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/buildTypes/id:%s", buildConfID)
	return c.doRetryRequest("DELETE", path, nil, nil)
}
