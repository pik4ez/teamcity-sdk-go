package types

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Display int
type ReadOnly bool

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

func (s ReadOnly) String() string {
	if s == true {
		return "true"
	}
	return "false"
}

func parseReadOnly(s string) ReadOnly {
	if s == "true" {
		return true
	}
	return false
}

type ParameterType interface {
	TypeName() string
	Values() map[string]string
}

type ParameterSpec struct {
	Label       string
	Description string
	Display     Display
	ReadOnly    ReadOnly
	Type        ParameterType
}

type Parameter struct {
	Value string
	Spec  *ParameterSpec
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
	Checked   string
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
	AllowMultiple  bool
	ValueSeparator string
	Items          []SelectItem
}

func (t SelectType) TypeName() string {
	return "select"
}

func (t SelectType) Values() map[string]string {
	ret := make(map[string]string)
	if t.AllowMultiple {
		ret["multiple"] = "true"
		ret["valueSeparator"] = t.ValueSeparator
	}
	for idx, item := range t.Items {
		ret[fmt.Sprintf("data_%d", idx+1)] = item.Value
		if item.Label != "" && item.Label != item.Value {
			ret[fmt.Sprintf("label_%d", idx+1)] = item.Label
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

func parseParameterType(s string, v map[string]string) ParameterType {
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
					itemMap[idx-1] = SelectItem{
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
			AllowMultiple:  v["multiple"] == "true",
			ValueSeparator: v["valueSeparator"],
			Items:          items,
		}
	}
	if s == "checkbox" {
		return CheckboxType{
			Checked:   v["checkedValue"],
			Unchecked: v["uncheckedValue"],
		}
	}
	return PasswordType{}
}

type Parameters map[string]Parameter

type parametersInput struct {
	Parameter []oneParameter `json:"property"`
}

type oneParameterType struct {
	RawValue string `json:"rawValue"`
}

type oneParameter struct {
	Name  string            `json:"name"`
	Value string            `json:"value"`
	Type  *oneParameterType `json:"type,omitempty"`
}

func (p Parameter) rawTypeValue() *oneParameterType {
	if p.Spec == nil {
		return nil
	}
	s := p.Spec
	return &oneParameterType{
		RawValue: s.String(),
	}
}

func (s ParameterSpec) String() string {
	typeText := s.Type.TypeName()
	values := s.Type.Values()
	if s.Label != "" {
		values["label"] = s.Label
	}
	if s.Description != "" {
		values["description"] = s.Description
	}
	values["display"] = s.Display.String()
	values["readOnly"] = s.ReadOnly.String()

	// Sort keys so that raw text is deterministic and testable
	valueKeys := make([]string, 0)
	for name := range values {
		valueKeys = append(valueKeys, name)
	}
	sort.Strings(valueKeys)

	valuesText := typeText
	for _, name := range valueKeys {
		value := values[name]
		valuesText += fmt.Sprintf(" %s='%s'", name,
			strings.Replace(strings.Replace(value, "|", "||", -1), "'", "|'", -1))
	}
	log.Printf("[DEBUG] Raw parameter spec %q\n", valuesText)
	return valuesText
}

var keyValue = regexp.MustCompile("\\w+='([^'|]|\\|\\||\\|')*'")

func (opt *oneParameterType) parseRawValue() *ParameterSpec {
	if opt == nil {
		return nil
	}
	sp := strings.SplitN(opt.RawValue, " ", 2)
	values := keyValue.FindAllString(sp[1], -1)
	specValue := make(map[string]string)
	for _, value := range values {
		kv := strings.SplitN(value, "=", 2)
		aValue := strings.Replace(strings.Replace(kv[1][1:len(kv[1])-1], "|'", "'", -1), "||", "|", -1)
		specValue[kv[0]] = aValue
	}
	return &ParameterSpec{
		Label:       specValue["label"],
		Description: specValue["description"],
		Display:     parseDisplay(specValue["display"]),
		ReadOnly:    parseReadOnly(specValue["readOnly"]),
		Type:        parseParameterType(sp[0], specValue),
	}
}

func (p Parameters) MarshalJSON() ([]byte, error) {
	pi := &parametersInput{
		Parameter: make([]oneParameter, 0),
	}
	for name, parameter := range p {
		pi.Parameter = append(pi.Parameter, oneParameter{
			Name:  name,
			Value: parameter.Value,
			Type:  parameter.rawTypeValue(),
		})
	}
	return json.Marshal(pi)
}

func (p *Parameters) UnmarshalJSON(b []byte) error {
	var pi parametersInput
	if err := json.Unmarshal(b, &pi); err != nil {
		return err
	}
	m := make(Parameters)
	for _, prop := range pi.Parameter {
		spec := prop.Type.parseRawValue()
		/*
		   pvt := PasswordType{}
		   if prop.Value == "" && spec != nil && spec.Type == pvt {
		      prop.Value = fmt.Sprintf("%%secure:teamcity.password.%s%%", prop.Name)
		   }
		*/
		m[prop.Name] = Parameter{
			Value: prop.Value,
			Spec:  spec,
		}
	}
	*p = m
	return nil
}

type NamedParameter struct {
	Name string
	Parameter
}

func (p NamedParameter) MarshalJSON() ([]byte, error) {
	pi := &oneParameter{
		Name:  p.Name,
		Value: p.Value,
		Type:  p.rawTypeValue(),
	}
	return json.Marshal(pi)
}
