package teamcity

import (
	"bytes"
	"fmt"
)

func (c *Client) SetProjectDescription(projectID, description string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/projects/id:%s/description", projectID)

	body := bytes.NewBuffer([]byte(description))
	_, err := c.doNotJSONRequest("PUT", path, "text/plain", "text/plain", body)
	if err != nil {
		return err
	}
	return nil
}
