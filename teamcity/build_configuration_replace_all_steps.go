package teamcity

import (
	"errors"
	"fmt"
	"github.com/Cardfree/teamcity-sdk-go/types"
)

func (c *Client) ReplaceAllBuildConfigurationSteps(buildConfID string, steps *types.BuildSteps) error {
	path := fmt.Sprintf("/httpAuth/app/rest/buildTypes/id:%s/steps", buildConfID)
	var buildstepsReturn *types.BuildSteps

	err := c.doRetryRequest("PUT", path, steps, &buildstepsReturn)
	if err != nil {
		return err
	}

	if buildstepsReturn == nil {
		return errors.New("build configuration steps not updated")
	}
	*steps = *buildstepsReturn

	return nil
}
