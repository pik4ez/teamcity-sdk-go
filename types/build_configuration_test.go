package types

import (
  "encoding/json"
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestTemplateIdParsing(t *testing.T) {
  var v TemplateId

  err := json.Unmarshal([]byte("null"), &v)
  assert.NoError(t, err)
  assert.Equal(t, TemplateId(""), v)

  err = json.Unmarshal([]byte("{\"id\":\"Tempy\"}"), &v)
  assert.NoError(t, err)
  assert.Equal(t, TemplateId("Tempy"), v)
}

func TestTemplateIdWriting(t *testing.T) {
  var v TemplateId
  v = ""
  b, err := json.Marshal(v)
  assert.NoError(t, err)
  assert.Equal(t, "null", string(b))

  v = "Tempy"
  b, err = json.Marshal(v)
  assert.NoError(t, err)
  assert.Equal(t, "{\"id\":\"Tempy\"}", string(v))
}
