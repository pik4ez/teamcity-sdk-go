package teamcity

import (
	"fmt"
)

func (c *Client) DeleteVcsRoot(VcsRootId string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/vcs-roots/id:%s", VcsRootId)
	return c.doRetryRequest("DELETE", path, nil, nil)
}
