package teamcity

import (
	"errors"
	"fmt"
	"github.com/Cardfree/teamcity-sdk-go/types"
)

func (c *Client) ReplaceAllBuildConfigurationParameters(buildConfID string, parameters *types.Parameters) error {
	path := fmt.Sprintf("/httpAuth/app/rest/buildTypes/id:%s/parameters", buildConfID)
	var parametersReturn *types.Parameters

	err := c.doRetryRequest("PUT", path, parameters, &parametersReturn)
	if err != nil {
		return err
	}

	if parametersReturn == nil {
		return errors.New("build configuration not updated")
	}
	*parameters = *parametersReturn

	return nil
}
