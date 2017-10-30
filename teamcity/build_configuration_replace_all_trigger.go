package teamcity

import (
	"errors"
	"fmt"
	"github.com/Cardfree/teamcity-sdk-go/types"
)

func (c *Client) ReplaceAllBuildConfigurationTriggers(buildConfID string, triggers *types.BuildTriggers) error {
	path := fmt.Sprintf("/httpAuth/app/rest/buildTypes/id:%s/triggers", buildConfID)
	var buildTriggersReturn *types.BuildTriggers

	err := c.doRetryRequest("PUT", path, triggers, &buildTriggersReturn)
	if err != nil {
		return err
	}

	if buildTriggersReturn == nil {
		return errors.New("build configuration triggers not updated")
	}
	*triggers = *buildTriggersReturn

	return nil
}
