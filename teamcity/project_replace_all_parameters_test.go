package teamcity

import (
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/require"
  "github.com/umweltdk/teamcity/types"
  "testing"
)

func TestClientReplaceAllProjectParameters(t *testing.T) {
  client, err := NewRealTestClient(t)
  require.NoError(t, err, "Expected no error")
  err = client.DeleteProject("TestClientReplaceAllProjectParameters")
  require.NoError(t, err, "Expected no error")
  err = client.CreateProject(&types.Project{
    Name: "TestClientReplaceAllProjectParameters",
  })
  require.NoError(t, err, "Expected no error")
  err = client.ReplaceAllProjectParameters("TestClientReplaceAllProjectParameters", &types.Properties{
    "env.MUH": types.Property{
      Value: "Hello",
      Spec: &types.PropertySpec{
        Type: types.PasswordType{},
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
  err = client.ReplaceAllProjectParameters("TestClientReplaceAllProjectParameters", &parameters)
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

  config, err := client.GetProject("TestClientReplaceAllProjectParameters")
  require.NoError(t, err, "Expected no error")
  require.NotNil(t, config, "Get to return config")
  assert.Equal(t, expected, config.Parameters)
}

func TestClientReplaceProjectParametersInherited(t *testing.T) {
  client, err := NewRealTestClient(t)
  require.NoError(t, err, "Expected no error")
  err = client.DeleteProject("Empty3")
  require.NoError(t, err, "Expected no error")
  err = client.CreateProject(&types.Project{
    Name:            "Empty3",
  })
  require.NoError(t, err, "Expected no error")
  err = client.ReplaceAllProjectParameters("Empty3", &types.Properties{
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
  err = client.DeleteProject("Empty3_TestClientReplaceProjectParametersInherited")
  require.NoError(t, err, "Expected no error")
  err = client.CreateProject(&types.Project{
    ID: "Empty3_TestClientReplaceProjectParametersInherited",
    ParentProjectID: "Empty3",
    Name: "TestClientReplaceProjectParametersInherited",
  })
  require.NoError(t, err, "Expected no error")
  err = client.ReplaceAllProjectParameters("Empty3_TestClientReplaceProjectParametersInherited", &types.Properties{
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
  err = client.ReplaceAllProjectParameters("Empty3_TestClientReplaceProjectParametersInherited", &parameters)
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

  config, err := client.GetProject("Empty3_TestClientReplaceProjectParametersInherited")
  require.NoError(t, err, "Expected no error")
  require.NotNil(t, config, "Get to return config")
  assert.Equal(t, expected, config.Parameters)
}
