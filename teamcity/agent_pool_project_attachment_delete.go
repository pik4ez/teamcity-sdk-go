package teamcity

import (
	"fmt"
)

func (c *Client) DeleteAgentPoolProjectAttachement(pool int, project string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/agentPools/id:%d/projects/%s", c.version, pool, project)
	return c.doRetryRequest("DELETE", path, nil, nil)
}
