package teamcity

import (
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/require"
  "github.com/umweltdk/teamcity/types"
  "testing"
  "time"
)

func TestClientCreateProjectMock(t *testing.T) {
  client := NewTestClient(newResponse(`{"id": "Empty_Plink", "parentProjectId":"Empty","name":"Plink"}`), nil)

  project := &types.Project{
    ParentProjectID: "Empty",
    Name:            "Plink",
  }

  err := client.CreateProject(project)
  require.NoError(t, err, "Expected no error")

  assert.Equal(t, "Empty_Plink", project.ID, "Expected create to return ID")
}

func TestClientCreateProjectMinimal(t *testing.T) {
  client, err := NewRealTestClient(t)
  require.NoError(t, err, "Expected no error")
  err = client.DeleteProject("Empty_Hello")
  require.NoError(t, err, "Expected no error")
  time.Sleep(5 * time.Second)

  project := &types.Project{
    ParentProjectID: "Empty",
    Name:            "Hello",
  }
  err = client.CreateProject(project)
  require.NoError(t, err, "Expected no error")
  require.NotNil(t, project, "Create to return project")

  assert.Equal(t, "Empty_Hello", project.ID, "Expected create to return ID")
  assert.Equal(t, types.Properties{
    "env.MUH": types.Property{
      Value: client.VersionParameterValue(t, "env.MUH"),
      Spec: &types.PropertySpec{
        Label: "Muh value",
        Description: "The Muh value that does all the Muhing",
        Display: types.Normal,
        Type: types.PasswordType{},
      },
    },
  }, project.Parameters, "no parameters")
}

func TestClientCreateProjectIgnoresParameters(t *testing.T) {
  client, err := NewRealTestClient(t)
  require.NoError(t, err, "Expected no error")
  err = client.DeleteProject("Empty_Full")
  require.NoError(t, err, "Expected no error")
  time.Sleep(5 * time.Second)

  project := &types.Project{
    ParentProjectID: "Empty",
    Name:            "Full",
    Parameters: types.Properties{
      "env.AWW": types.Property {
        Value: "Parent",
      },
    },
  }
  err = client.CreateProject(project)
  require.NoError(t, err, "Expected no error")
  require.NotNil(t, project, "Create to return project")

  assert.Equal(t, "Empty_Full", project.ID, "Expected create to return ID")
  assert.Equal(t, types.Properties{
    "env.MUH": types.Property{
      Value: client.VersionParameterValue(t, "env.MUH"),
      Spec: &types.PropertySpec{
        Label: "Muh value",
        Description: "The Muh value that does all the Muhing",
        Display: types.Normal,
        Type: types.PasswordType{},
      },
    },
  }, project.Parameters, "no parameters")
}

func TestClientCreateProjectUsedID(t *testing.T) {
  client, err := NewRealTestClient(t)
  require.NoError(t, err, "Expected no error")
  client.retries = 1

  project := &types.Project{
    ID:        "Single",
    Name:      "Hej Med Dig",
  }

  err = client.CreateProject(project)
  assert.Error(t, err, "Expected error")
}

func TestClientCreateProjectUsedName(t *testing.T) {
  client, err := NewRealTestClient(t)
  require.NoError(t, err, "Expected no error")
  client.retries = 1

  project := &types.Project{
    Name:      "Single",
  }

  err = client.CreateProject(project)
  assert.Error(t, err, "Expected error")
}

func TestClientCreateProjectUsedNameExplicitID(t *testing.T) {
  client, err := NewRealTestClient(t)
  require.NoError(t, err, "Expected no error")
  client.retries = 1

  project := &types.Project{
    ID:        "Single_Dubie",
    Name:      "Single",
  }

  err = client.CreateProject(project)
  assert.Error(t, err, "Expected error")
}