package teamcity

import (
	"fmt"
)

func (c *Client) DeleteBuildConfigurationSetting(buildConfID, name string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/buildTypes/id:%s/settings/%s", c.version, buildConfID, name)
	return c.doRetryRequest("DELETE", path, nil, nil)
}
