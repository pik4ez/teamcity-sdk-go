package teamcity

import (
	"errors"
	"fmt"
	"github.com/Cardfree/teamcity-sdk-go/types"
)

func (c *Client) ReplaceAllBuildConfigurationAgentRequirements(buildConfID string, agentRequirements *types.BuildAgentRequirements) error {
	path := fmt.Sprintf("/httpAuth/app/rest/buildTypes/id:%s/agent-requirements", buildConfID)
	var buildAgentRequirementsReturn *types.BuildAgentRequirements

	err := c.doRetryRequest("PUT", path, agentRequirements, &buildAgentRequirementsReturn)
	if err != nil {
		return err
	}

	if buildAgentRequirementsReturn == nil {
		return errors.New("build configuration agent requirements not updated")
	}
	*agentRequirements = *buildAgentRequirementsReturn

	return nil
}
