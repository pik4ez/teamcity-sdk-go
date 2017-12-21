package teamcity

import (
	"fmt"
)

func (c *Client) DeleteVcsRoot(VcsRootId string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/vcs-roots/id:%s", c.version, VcsRootId)
	return c.doRetryRequest("DELETE", path, nil, nil)
}
