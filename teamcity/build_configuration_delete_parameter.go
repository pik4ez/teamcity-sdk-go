package teamcity

import (
	"fmt"
)

func (c *Client) DeleteBuildConfigurationParameter(buildConfID, name string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/buildTypes/id:%s/parameters/%s", buildConfID, name)
	return c.doRetryRequest("DELETE", path, nil, nil)
}
