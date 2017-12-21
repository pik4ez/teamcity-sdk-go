package teamcity

import (
	"bytes"
	"fmt"
)

func (c *Client) ReplaceBuildConfigurationParameterValue(buildConfID, name string, value string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/buildTypes/id:%s/parameters/%s", c.version, buildConfID, name)

	body := bytes.NewBuffer([]byte(value))
	_, err := c.doNotJSONRequest("PUT", path, "text/plain", "text/plain", body)
	if err != nil {
		return err
	}
	return nil
}
