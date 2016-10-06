package types

import (
  "encoding/json"
  "github.com/stretchr/testify/assert"
  //"github.com/stretchr/testify/require"
  "testing"
)

func MarshalAndUnmarhalMatch(t *testing.T, jsonValue string, props *Properties) {
  bytes, err := json.Marshal(&props)
  assert.NoError(t, err)
  assert.Equal(t, jsonValue, string(bytes))

  var val *Properties
  err = json.Unmarshal([]byte(jsonValue), &val)
  assert.NoError(t, err)
  assert.Equal(t, props, val)
}

func TestProperties(t *testing.T) {
  MarshalAndUnmarhalMatch(t, `{"property":[{"name":"test","value":"muh"}]}`,
    &Properties{
      "test": Property{
        Value: "muh",
      },
    })

  MarshalAndUnmarhalMatch(t, `{"property":[{"name":"env.TEST_RUNNER","value":"l","type":{"rawValue":"password description='What test runner are we going to use' display='normal' label='Test runner'"}}]}`,
    &Properties{
      "env.TEST_RUNNER": Property{
        Value: "l",
        Spec: &PropertySpec{
          Type: PasswordType{},
          Label: "Test runner",
          Description: "What test runner are we going to use",
        },
      },
    })

  MarshalAndUnmarhalMatch(t, `{"property":[{"name":"env.TEST_RUNNER","value":"l","type":{"rawValue":"text description='What test runner are we going to use' display='prompt' validationMode='not_empty'"}}]}`,
    &Properties{
      "env.TEST_RUNNER": Property{
        Value: "l",
        Spec: &PropertySpec{
          Type: TextType{"not_empty"},
          Description: "What test runner are we going to use",
          Display: Prompt,
        },
      },
    })

  MarshalAndUnmarhalMatch(t, `{"property":[{"name":"env.TEST_RUNNER","value":"l","type":{"rawValue":"checkbox checkedValue='Wow' display='hidden' label='Test runner'"}}]}`,
    &Properties{
      "env.TEST_RUNNER": Property{
        Value: "l",
        Spec: &PropertySpec{
          Type: CheckboxType{Checked: "Wow"},
          Label: "Test runner",
          Display: Hidden,
        },
      },
    })

  MarshalAndUnmarhalMatch(t, `{"property":[{"name":"env.PLING","value":"donk","type":{"rawValue":"select data_1='te' display='hidden' label='Test |||| ||M |||'runner||'"}}]}`,
    &Properties{
      "env.PLING": Property{
        Value: "donk",
        Spec: &PropertySpec{
          Type: SelectType{Items:[]SelectItem {
            SelectItem{"","te"},
          }},
          Label: "Test || |M |'runner|",
          Display: Hidden,
        },
      },
    })
}
