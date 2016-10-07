package teamcity

import (
  "errors"
  "fmt"
  "github.com/umweltdk/teamcity/types"
)

func (c *Client) ReplaceProjectParameter(projectID,name string, parameter *types.Parameter) error {
  path := fmt.Sprintf("/httpAuth/app/rest/projects/id:%s/parameters/%s", projectID, name)
  var parameterReturn *types.Parameter

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
