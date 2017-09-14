package teamcity

import (
	"errors"
	"github.com/Cardfree/teamcity-sdk-go/types"
)

func (c *Client) CreateVcsRoot(vcs *types.VcsRoot) error {
	path := "/httpAuth/app/rest/vcs-roots"
	var vcsReturn *types.VcsRoot

	err := c.doRetryRequest("POST", path, vcs, &vcsReturn)
	if err != nil {
		return err
	}

	if vcsReturn == nil {
		return errors.New("VCS Root not created")
	}
	*vcs = *vcsReturn

	return nil
}
