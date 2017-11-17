package teamcity

import (
	"fmt"
)

func (c *Client) DeleteAgentPoolAttachement(pool string, project string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/agentPools/name:%s/projects/%s", pool, project)
	return c.doRetryRequest("DELETE", path, nil, nil)
}
