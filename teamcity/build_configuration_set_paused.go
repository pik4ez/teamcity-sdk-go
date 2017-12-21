package teamcity

import (
	"bytes"
	"fmt"
	"strconv"
)

func (c *Client) SetBuildConfigurationPaused(buildConfID, state bool) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/buildTypes/id:%s/paused", c.version, buildConfID)

	body := bytes.NewBuffer([]byte(strconv.FormatBool(state)))
	_, err := c.doNotJSONRequest("PUT", path, "text/plain", "text/plain", body)
	if err != nil {
		return err
	}
	return nil
}
