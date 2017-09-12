package teamcity

import (
	"github.com/Cardfree/teamcity-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	// "time"
)

func TestClientCreateVcsRootMock(t *testing.T) {
	client := NewTestClient(newResponse(`
	{
	  "id": "Empty_Plink",
	  "name":"Plink",
	  "VcsName":"jetbrains.git"
	}
	`), nil)

	vcs := &types.VcsRoot{
		ID:      "Empty",
		Name:    "Plink",
		VcsName: "jetbrains.git",
	}

	err := client.CreateVcsRoot(vcs)
	require.NoError(t, err, "Expected no error")

	assert.Equal(t, "Empty_Plink", vcs.ID, "Expected create to return ID")
	assert.Equal(t, "Plink", vcs.Name, "Expected create to return Name")
	assert.Equal(t, "jetbrains.git", vcs.VcsName, "Expected create to return VcsName")

}

func TestClientCreateVcsRootMinimal(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteVcsRoot("Empty_Plink")
	require.NoError(t, err, "Expected no error")

	vcs := &types.VcsRoot{
		ID:        "Empty_Plink",
		Name:      "Plink",
		VcsName:   "jetbrains.git",
		ProjectID: "Empty",

		Properties: types.Properties{
			"url":    "https://github.com/cardfree/teamcity-sdk-go",
			"branch": "refs/heads/master",
		},
	}

	err = client.CreateVcsRoot(vcs)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, vcs, "Create to return project")

	assert.Equal(t, "Empty_Plink", vcs.ID, "Expected create to return ID")
	// assert.Equal(t, types.ProjectId("Empty"), vcs.ProjectID, "Expected create to return ParentProjectID")
	assert.Equal(t, "Plink", vcs.Name, "Expected create to return Name")
	assert.Equal(t, "jetbrains.git", vcs.VcsName, "Expected create to return VcsName")
	assert.Equal(t, types.Properties{
		"url":    "https://github.com/cardfree/teamcity-sdk-go",
		"branch": "refs/heads/master",
	}, vcs.Properties, "no properties")
}

func TestClientCreateVcsRootUsedID(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	client.retries = 1

	vcs := &types.VcsRoot{
		ID:        "Empty_Plink",
		Name:      "Dlink",
		ProjectID: "Empty",
	}

	err = client.CreateVcsRoot(vcs)
	assert.Error(t, err, "Expected error")
}
