package teamcity

import (
	"fmt"
	"github.com/Cardfree/teamcity-sdk-go/types"
)

func (c *Client) GetProject(projectID string) (*types.Project, error) {
	path := fmt.Sprintf("/httpAuth/app/rest/projects/id:%s", projectID)
	var project *types.Project

	err := c.doRetryRequest("GET", path, nil, &project)
	if err != nil {
		return nil, err
	}

	return project, nil
}
