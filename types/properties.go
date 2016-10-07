package types

import (
	"encoding/json"
)

type Properties map[string]string

type propertiesInput struct {
	Property []oneProperty `json:"property"`
}

type oneProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (p Properties) MarshalJSON() ([]byte, error) {
	pi := &propertiesInput{
		Property: make([]oneProperty, 0),
	}
	for name, property := range p {
		pi.Property = append(pi.Property, oneProperty{
			Name:  name,
			Value: property,
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
		m[prop.Name] = prop.Value
	}
	*p = m
	return nil
}
