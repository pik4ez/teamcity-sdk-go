package teamcity

import (
	"errors"
	"fmt"
	"github.com/Cardfree/teamcity-sdk-go/types"
)

func (c *Client) ReplaceAllBuildConfigurationFeatures(buildConfID string, features *types.BuildFeatures) error {
	path := fmt.Sprintf("/httpAuth/app/rest/buildTypes/id:%s/features", buildConfID)
	var buildFeaturesReturn *types.BuildFeatures

	err := c.doRetryRequest("PUT", path, features, &buildFeaturesReturn)
	if err != nil {
		return err
	}

	if buildFeaturesReturn == nil {
		return errors.New("build configuration features not updated")
	}
	*features = *buildFeaturesReturn

	return nil
}
