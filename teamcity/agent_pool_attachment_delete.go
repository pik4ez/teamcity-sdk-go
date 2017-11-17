package teamcity

import (
	"fmt"
)

func (c *Client) DeleteAgentPoolAttachement(name string, project string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/agentPools/name:%s/%s", name, project)
	return c.doRetryRequest("DELETE", path, nil, nil)
}
