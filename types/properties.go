package types

import (
	"encoding/json"
  "fmt"
  "regexp"
  "sort"
  "strings"
  "strconv"
)

type Display int

const (
    Normal Display = iota
    Hidden
    Prompt
)

func (s Display) String() string {
  if s == Normal {
    return "normal"
  }
  if s == Hidden {
    return "hidden"
  }
  return "prompt"
}

func parseDisplay(s string) Display {
  if s == "hidden" {
    return Hidden
  }
  if s == "prompt" {
    return Prompt
  }
  return Normal
}

type PropertyType interface {
    TypeName() string
    Values() map[string]string
}

type PropertySpec struct {
  Label string
  Description string
  Display Display
  Type PropertyType
}

type Property struct {
  Value string
  Spec *PropertySpec
}


type PasswordType struct {
}

func (t PasswordType) TypeName() string {
  return "password"
}

func (t PasswordType) Values() map[string]string {
  return make(map[string]string)
}


type CheckboxType struct {
  Checked string
  Unchecked string
}

func (t CheckboxType) TypeName() string {
  return "checkbox"
}

func (t CheckboxType) Values() map[string]string {
  ret := make(map[string]string)
  if t.Checked != "" {
    ret["checkedValue"] = t.Checked 
  }
  if t.Unchecked != "" {
    ret["uncheckedValue"] = t.Unchecked
  }
  return ret
}


type SelectItem struct {
  Label string
  Value string
}

type SelectType struct {
  AllowMutiple bool
  ValueSeparator string
  Items []SelectItem
}

func (t SelectType) TypeName() string {
  return "select"
}

func (t SelectType) Values() map[string]string {
  ret := make(map[string]string)
  if t.AllowMutiple {
    ret["multiple"] = "true"
    ret["valueSeparator"] = t.ValueSeparator 
  }
  for idx, item := range t.Items {
    ret[fmt.Sprintf("data_%d", idx + 1)] = item.Value
    if item.Label != "" && item.Label != item.Value {
      ret[fmt.Sprintf("label_%d", idx + 1)] = item.Label
    }
  }
  return ret
}


type TextType struct {
  ValidationMode string
}

func (t TextType) TypeName() string {
  return "text"
}

func (t TextType) Values() map[string]string {
  ret := make(map[string]string)
  ret["validationMode"] = t.ValidationMode 
  return ret
}

var dataRegex = regexp.MustCompile("^data_(\\d+)$")

func parsePropertyType(s string, v map[string]string) PropertyType {
  if s == "text" {
    return TextType{
      ValidationMode: v["validationMode"],
    }
  }
  if s == "select" {
    length := 0
    itemMap := make(map[int]SelectItem)
    for name, value := range v {
      p := dataRegex.FindStringSubmatch(name)
      if p != nil {
        idx, err := strconv.Atoi(p[1])
        if err == nil {
          itemMap[idx - 1] = SelectItem{
            Value: value,
            Label: v[fmt.Sprintf("label_%d", idx)],
          }
          if idx > length {
            length = idx
          }
        }
      }
    }
    items := make([]SelectItem, length)
    for idx := range items {
      items[idx] = itemMap[idx]
    }
    return SelectType{
      AllowMutiple: v["multiple"] == "true",
      ValueSeparator: v["valueSeparator"],
      Items: items,
    }
  }
  if s == "checkbox" {
    return CheckboxType{
      Checked: v["checkedValue"],
      Unchecked: v["uncheckedValue"],
    }
  }
  return PasswordType{}
}


type Properties map[string]Property

type propertiesInput struct {
	Property []oneProperty `json:"property"`
}

type onePropertyType struct {
  RawValue string `json:"rawValue"`
}

type oneProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
  Type  *onePropertyType `json:"type,omitempty"`
}

func(p Property) rawTypeValue() *onePropertyType {
  if p.Spec == nil {
    return nil
  }
  s := p.Spec
  typeText := s.Type.TypeName()
  values := s.Type.Values()
  if s.Label != "" {
    values["label"] = s.Label
  }
  if s.Description != "" {
    values["description"] = s.Description
  }
  values["display"] = s.Display.String()

  // Sort keys so that raw text is deterministic and testable
  valueKeys := make([]string, 0)
  for name, _ := range values {
    valueKeys = append(valueKeys, name)
  }
  sort.Strings(valueKeys)

  valuesText := typeText
  for _, name := range valueKeys {
    value := values[name]
    valuesText += fmt.Sprintf(" %s='%s'", name, 
      strings.Replace(strings.Replace(value, "|", "||", -1), "'", "|'", -1))
  }
  return &onePropertyType{
    RawValue: valuesText,
  }
}

var keyValue = regexp.MustCompile("\\w+='([^'|]|\\|\\||\\|')*'")

func (opt *onePropertyType) parseRawValue() *PropertySpec {
  if opt == nil {
    return nil
  }
  sp := strings.SplitN(opt.RawValue, " ", 2)
  values := keyValue.FindAllString(sp[1], -1)
  specValue := make(map[string]string)
  for _, value := range values {
    kv := strings.SplitN(value, "=", 2)
    aValue := strings.Replace(strings.Replace(kv[1][1:len(kv[1]) - 1], "|'", "'", -1), "||", "|", -1)
    specValue[kv[0]] = aValue
  }
  return &PropertySpec{
    Label: specValue["label"],
    Description: specValue["description"],
    Display: parseDisplay(specValue["display"]),
    Type: parsePropertyType(sp[0], specValue),
  }
}

func (p Properties) MarshalJSON() ([]byte, error) {
	pi := &propertiesInput{
		Property: make([]oneProperty, 0),
	}
	for name, property := range p {
		pi.Property = append(pi.Property, oneProperty{
			Name:  name,
			Value: property.Value,
      Type: property.rawTypeValue(),
		})
	}
	return json.Marshal(pi)
}

func (p *Properties) UnmarshalJSON(b []byte) error {
	var pi propertiesInput
	if err := json.Unmarshal(b, &pi); err != nil {
		return err
	}
	m := make(Properties)
	for _, prop := range pi.Property {
    spec := prop.Type.parseRawValue()
    /*
    pvt := PasswordType{}
    if prop.Value == "" && spec != nil && spec.Type == pvt {
       prop.Value = fmt.Sprintf("%%secure:teamcity.password.%s%%", prop.Name)
    }
    */
		m[prop.Name] = Property{
      Value: prop.Value,
      Spec: spec,
    }
	}
	*p = m
	return nil
}
