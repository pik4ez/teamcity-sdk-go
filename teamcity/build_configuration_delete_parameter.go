package teamcity

import (
	"fmt"
)

func (c *Client) DeleteBuildConfigurationParameter(buildConfID, name string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/buildTypes/id:%s/parameters/%s", c.version, buildConfID, name)
	return c.doRetryRequest("DELETE", path, nil, nil)
}
