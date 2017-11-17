package teamcity

import (
	"fmt"
	"github.com/Cardfree/teamcity-sdk-go/types"
)

func (c *Client) GetAgentPool(name string) (*types.AgentPools, error) {
	path := fmt.Sprintf("/httpAuth/app/rest/agentPools/name:%s", name)
	var agp *types.AgentPools

	err := c.doRetryRequest("GET", path, nil, &agp)
	if err != nil {
		return nil, err
	}

	return agp, nil
}
