package teamcity

import (
  "errors"
  "fmt"
  "github.com/umweltdk/teamcity/types"
)

func (c *Client) ReplaceBuildConfigurationParameter(buildConfID,name string, parameter *types.Property) error {
  path := fmt.Sprintf("/httpAuth/app/rest/buildTypes/id:%s/parameters/%s", buildConfID, name)
  var parameterReturn *types.Property

  err := c.doRetryRequest("PUT", path, parameter, &parameterReturn)
  if err != nil {
    return err
  }

  if parameterReturn == nil {
    return errors.New("parameter not updated")
  }
  *parameter = *parameterReturn

  return nil
}
