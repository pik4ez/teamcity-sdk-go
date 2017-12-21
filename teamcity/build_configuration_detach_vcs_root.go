package teamcity

import (
	"fmt"
)

func (c *Client) DetachBuildConfigurationVcsRoot(buildConfID string, vcsRootID string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/buildTypes/id:%s/vcs-root-entries/%s", c.version, buildConfID, vcsRootID)
	return c.doRetryRequest("DELETE", path, nil, nil)
}
