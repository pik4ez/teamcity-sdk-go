package teamcity

import (
	"bytes"
	"fmt"
)

func (c *Client) SetBuildConfigurationDescription(buildConfID, description string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/buildTypes/id:%s/description", buildConfID)

	body := bytes.NewBuffer([]byte(description))
	_, err := c.doNotJSONRequest("PUT", path, "text/plain", "text/plain", body)
	if err != nil {
		return err
	}
	return nil
}
