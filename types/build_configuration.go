package types

import (
	"encoding/json"
)

type BuildConfiguration struct {
	ID             string         `json:"id,omitempty"`
	ProjectID      string         `json:"projectId"`
	TemplateFlag   bool           `json:"templateFlag"`
	Template       *TemplateID    `json:"template,omitempty"`
	Name           string         `json:"name"`
	Description    string         `json:"description,omitempty"`
	Settings       Properties     `json:"settings,omitempty"`
	Parameters     Parameters     `json:"parameters,omitempty"`
	Steps          BuildSteps     `json:"steps,omitempty"`
	VcsRootEntries VcsRootEntries `json:"vcs-root-entries,omitempty"`
	//Features []Feature
	//Triggers []Trigger
	//AgentRequirements []AgentRequirement
}

type BuildConfigurations map[string]BuildConfiguration

type buildConfigurationsInput struct {
	BuildType []BuildConfiguration
}

func (bc BuildConfigurations) MarshalJSON() ([]byte, error) {
	bci := &buildConfigurationsInput{
		BuildType: make([]BuildConfiguration, 0),
	}
	for _, value := range bc {
		bci.BuildType = append(bci.BuildType, value)
	}
	return json.Marshal(bci)
}

func (bc *BuildConfigurations) UnmarshalJSON(b []byte) error {
	var bci buildConfigurationsInput
	if err := json.Unmarshal(b, &bci); err != nil {
		return err
	}
	m := make(BuildConfigurations)
	for _, bt := range bci.BuildType {
		m[bt.ID] = bt
	}
	*bc = m
	return nil
}
