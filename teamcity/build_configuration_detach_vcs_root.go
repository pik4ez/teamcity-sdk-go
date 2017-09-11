package teamcity

import (
	"fmt"
)

func (c *Client) DetachBuildConfigurationVcsRoot(buildConfID string, vcsRootID string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/buildTypes/id:%s/vcs-root-entries/%s", buildConfID, vcsRootID)
	return c.doRetryRequest("DELETE", path, nil, nil)
}
