package teamcity

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) ProjectRename(projectID, name string) error {
	path := fmt.Sprintf("/app/rest/projects/id:%s/name", projectID)

	req, err := http.NewRequest("PUT", c.host+path, bytes.NewReader([]byte(name)))
	up := base64.RawStdEncoding.EncodeToString([]byte(c.username + ":" + c.password))
	req.Header.Add("Authorization", "Basic "+up)

	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
