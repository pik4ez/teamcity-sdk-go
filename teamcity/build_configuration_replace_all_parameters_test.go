package teamcity

import (
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/require"
  "github.com/umweltdk/teamcity/types"
  "testing"
)

func TestClientAllReplaceBuildConfigurationParameters(t *testing.T) {
  client, err := NewRealTestClient(t)
  require.NoError(t, err, "Expected no error")
  err = client.DeleteBuildConfiguration("Single_TestClientUpdateBuildConfiguration")
  require.NoError(t, err, "Expected no error")
  err = client.CreateBuildConfiguration(&types.BuildConfiguration{
    ID: "Single_TestClientUpdateBuildConfiguration",
    ProjectID: "Single",
    Name: "Test Client Update Build Configuration",
    Parameters: types.Properties{
      "env.MUH": types.Property{
        Value: "Hello",
        Spec: &types.PropertySpec{
          Type: types.PasswordType{},
        },
      },
    },
  })
  require.NoError(t, err, "Expected no error")

  parameters := types.Properties{
    "env.HELLO": types.Property{"Good job", nil},
    "aws.hush": types.Property{
      Value: "Bad Job",
      Spec: &types.PropertySpec{
        Type: types.PasswordType{},
      },
    },
  }
  err = client.ReplaceAllBuildConfigurationParameters("Single_TestClientUpdateBuildConfiguration", &parameters)
  require.NoError(t, err, "Expected no error")
  require.NotNil(t, parameters, "Update to return parameters")

  expected := types.Properties{
    "env.HELLO": types.Property{"Good job", nil},
    "aws.hush": types.Property{
      Value: "",
      Spec: &types.PropertySpec{
        Type: types.PasswordType{},
      },
    },
  }
  assert.Equal(t, expected, parameters)

  config, err := client.GetBuildConfiguration("Single_TestClientUpdateBuildConfiguration")
  require.NoError(t, err, "Expected no error")
  require.NotNil(t, config, "Get to return config")
  assert.Equal(t, expected, config.Parameters)
}

func TestClientReplaceAllBuildConfigurationParametersInherited(t *testing.T) {
  client, err := NewRealTestClient(t)
  require.NoError(t, err, "Expected no error")
  err = client.DeleteProject("Empty2")
  require.NoError(t, err, "Expected no error")
  err = client.CreateProject(&types.Project{
    ParentProjectID: "",
    Name:            "Empty2",
  })
  require.NoError(t, err, "Expected no error")
  err = client.ReplaceAllProjectParameters("Empty2", &types.Properties{
    "env.MUH": types.Property{
      Value: "Hush",
      Spec: &types.PropertySpec{
        Label: "Muh value",
        Description: "The Muh value that does all the Muhing",
        Type: types.PasswordType{},
      },
    },
    "env.AWW": types.Property{
      Value: "Parent",
      Spec: &types.PropertySpec{
        Label: "AWW",
        Type: types.TextType{"any"},
      },
    },
    "env.DAMM": types.Property{
      Value: "Parent",
    },
  })
  err = client.DeleteBuildConfiguration("Empty2_TestClientUpdateBuildConfigurationInherited")
  require.NoError(t, err, "Expected no error")
  err = client.CreateBuildConfiguration(&types.BuildConfiguration{
    ID: "Empty2_TestClientUpdateBuildConfigurationInherited",
    ProjectID: "Empty2",
    Name: "Test Client Update Build Configuration",
    Parameters: types.Properties{
      "config.remove": types.Property{
        Value: "Hello",
      },
      "config.replace": types.Property{
        Value: "Dink",
        Spec: &types.PropertySpec{
          Label: "Buhhhhh",
          Type: types.TextType{"any"},
        },
      },
      "env.AWW": types.Property{
        Value: "BuildConf",
      },
      "env.DAMM": types.Property{
        Value: "BuildConf",
      },
    },
  })
  require.NoError(t, err, "Expected no error")

  parameters := types.Properties{
    "env.HELLO": types.Property{"Good job", nil},
    "config.replace": types.Property{
      Value: "Mink",
      Spec: &types.PropertySpec{
        Label: "Minker",
        Type: types.CheckboxType{
          Checked: "Flunk",
        },
      },
    },
    "aws.hush": types.Property{
      Value: "Bad Job",
      Spec: &types.PropertySpec{
        Type: types.PasswordType{},
      },
    },
    "env.AWW": types.Property{
      Value: "",
      Spec: &types.PropertySpec{
        Type: types.CheckboxType{
          Checked: "BuildCD",
        },
      },
    },
    "env.MUH": types.Property{
      Value: "Hello",
      Spec: &types.PropertySpec{
        Label: "Plunk",
        Type: types.PasswordType{},
      },
    },
  }
  err = client.ReplaceAllBuildConfigurationParameters("Empty2_TestClientUpdateBuildConfigurationInherited", &parameters)
  require.NoError(t, err, "Expected no error")
  require.NotNil(t, parameters, "Update to return parameters")

  expected := types.Properties{
    "env.HELLO": types.Property{"Good job", nil},
    "config.replace": types.Property{
      Value: "Mink",
      Spec: &types.PropertySpec{
        Label: "Minker",
        Type: types.CheckboxType{
          Checked: "Flunk",
        },
      },
    },
    "aws.hush": types.Property{
      Value: "",
      Spec: &types.PropertySpec{
        Type: types.PasswordType{},
      },
    },
    "env.AWW": types.Property{
      Value: "",
      Spec: &types.PropertySpec{
        Label: "AWW",
        Type: types.TextType{"any"},
      },
    },
    "env.MUH": types.Property{
      Value: "",
      Spec: &types.PropertySpec{
        Label: "Muh value",
        Description: "The Muh value that does all the Muhing",
        Type: types.PasswordType{},
      },
    },
    "env.DAMM": types.Property{
      Value: "Parent",
    },
  }
  assert.Equal(t, expected, parameters)

  config, err := client.GetBuildConfiguration("Empty2_TestClientUpdateBuildConfigurationInherited")
  require.NoError(t, err, "Expected no error")
  require.NotNil(t, config, "Get to return config")
  assert.Equal(t, expected, config.Parameters)
}
