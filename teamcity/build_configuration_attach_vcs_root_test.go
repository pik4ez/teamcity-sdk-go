package teamcity

import (
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/require"
  "github.com/umweltdk/teamcity/types"
  "testing"
)

func TestClientAttachBuildConfigurationVcsRoot(t *testing.T) {
  client, err := NewRealTestClient(t)
  require.NoError(t, err, "Expected no error")
  err = client.DeleteBuildConfiguration("Single_TestClientAttachBuildConfigurationVcsRoot")
  require.NoError(t, err, "Expected no error")
  err = client.CreateBuildConfiguration(&types.BuildConfiguration{
    ID: "Single_TestClientAttachBuildConfigurationVcsRoot",
    ProjectID: "Single",
    Name: "TestClientAttachBuildConfigurationVcsRoot",
  })
  require.NoError(t, err, "Expected no error")

  vcsRoots := types.VcsRootEntries{
    types.VcsRootEntry{
      ID: "Flflfl",
      VcsRootID: "Single_HttpsGithubComUmweltdkDockerNodeGit",
      CheckoutRules: "+:refs/heads/master\n+:refs/heads/trigger*",
    },
  }
  err = client.AttachBuildConfigurationVcsRoot("Single_TestClientAttachBuildConfigurationVcsRoot", &vcsRoots[0])
  require.NoError(t, err, "Expected no error")
  require.NotNil(t, vcsRoots[0], "Update to return parameters")

  expected := types.VcsRootEntries{
    types.VcsRootEntry{
      ID: "Single_HttpsGithubComUmweltdkDockerNodeGit",
      VcsRootID: "Single_HttpsGithubComUmweltdkDockerNodeGit",
      CheckoutRules: "+:refs/heads/master\n+:refs/heads/trigger*",
    },
  }
  assert.Equal(t, expected, vcsRoots)

  config, err := client.GetBuildConfiguration("Single_TestClientAttachBuildConfigurationVcsRoot")
  require.NoError(t, err, "Expected no error")
  require.NotNil(t, config, "Get to return config")
  assert.Equal(t, expected, config.VcsRootEntries)
}

func TestClientAttachBuildConfigurationVcsRootExisting(t *testing.T) {
  client, err := NewRealTestClient(t)
  require.NoError(t, err, "Expected no error")
  err = client.DeleteBuildConfiguration("Single_TestClientAttachBuildConfigurationVcsRoot")
  require.NoError(t, err, "Expected no error")
  err = client.CreateBuildConfiguration(&types.BuildConfiguration{
    ID: "Single_TestClientAttachBuildConfigurationVcsRoot",
    ProjectID: "Single",
    Name: "TestClientAttachBuildConfigurationVcsRoot",
    VcsRootEntries: types.VcsRootEntries{
      types.VcsRootEntry{
        VcsRootID: "Single_HttpsGithubComUmweltdkDockerNodeGit",
        CheckoutRules: "+:refs/heads/master\n+:refs/heads/trigger*",
      },
    },
  })
  require.NoError(t, err, "Expected no error")

  vcsRoots := types.VcsRootEntries{
    types.VcsRootEntry{
      VcsRootID: "Single_HttpsGithubComUmweltdkDockerNodeGit",
    },
  }
  err = client.AttachBuildConfigurationVcsRoot("Single_TestClientAttachBuildConfigurationVcsRoot", &vcsRoots[0])
  require.NoError(t, err, "Expected no error")
  require.NotNil(t, vcsRoots[0], "Update to return parameters")

  expected := types.VcsRootEntries{
    types.VcsRootEntry{
      ID: "Single_HttpsGithubComUmweltdkDockerNodeGit",
      VcsRootID: "Single_HttpsGithubComUmweltdkDockerNodeGit",
      CheckoutRules: "",
    },
  }
  assert.Equal(t, expected, vcsRoots)

  config, err := client.GetBuildConfiguration("Single_TestClientAttachBuildConfigurationVcsRoot")
  require.NoError(t, err, "Expected no error")
  require.NotNil(t, config, "Get to return config")
  assert.Equal(t, expected, config.VcsRootEntries)
}

func TestClientAttachBuildConfigurationVcsRootWrongProject(t *testing.T) {
  client, err := NewRealTestClient(t)
  require.NoError(t, err, "Expected no error")
  err = client.DeleteBuildConfiguration("Empty_TestClientAttachBuildConfigurationVcsRoot")
  require.NoError(t, err, "Expected no error")
  err = client.CreateBuildConfiguration(&types.BuildConfiguration{
    ID: "Empty_TestClientAttachBuildConfigurationVcsRoot",
    ProjectID: "Empty",
    Name: "TestClientAttachBuildConfigurationVcsRoot",
  })
  require.NoError(t, err, "Expected no error")

  vcsRoots := types.VcsRootEntries{
    types.VcsRootEntry{
      VcsRootID: "Single_HttpsGithubComUmweltdkDockerNodeGit",
    },
  }
  err = client.AttachBuildConfigurationVcsRoot("Empty_TestClientAttachBuildConfigurationVcsRoot", &vcsRoots[0])
  assert.Error(t, err, "Expected error");
}