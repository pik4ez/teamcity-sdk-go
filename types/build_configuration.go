package types

import (
	"encoding/json"
)

type BuildConfiguration struct {
	ID                   string                    `json:"id,omitempty"`
	ProjectID            string                    `json:"projectId"`
	TemplateFlag         bool                      `json:"templateFlag"`
	TemplateID           TemplateId                `json:"template,omitempty"`
	Name                 string                    `json:"name"`
	Description          string                    `json:"description,omitempty"`
	VcsRootEntries       VcsRootEntries            `json:"vcs-root-entries,omitempty"`
	Settings             Properties                `json:"settings,omitempty"`
	Parameters           Parameters                `json:"parameters,omitempty"`
	Steps                BuildSteps                `json:"steps,omitempty"`
	Features             BuildFeatures             `json:"features,omitempty"`
	Triggers             BuildTriggers             `json:"triggers,omitempty"`
	SnapshotDependencies BuildSnapshotDependencies `json:"snapshot-dependencies,omitempty"`
	ArtifactDependencies BuildArtifactDependencies `json:"artifact-dependencies,omitempty"`
	AgentRequirements    BuildAgentRequirements    `json:"agent-requirements,omitempty"`
}

type TemplateId string

type BuildConfigurationShort struct {
	ID           string `json:"id"`
	ProjectID    string `json:"projectId,omitempty"`
	TemplateFlag bool   `json:"templateFlag,omitempty"`
	Name         string `json:"name,omitempty"`
	Href         string `json:"href,omitempty"`
}

func (p TemplateId) MarshalJSON() ([]byte, error) {
	if p != "" {
		pi := &BuildConfigurationShort{
			ID: string(p),
		}
		return json.Marshal(pi)
	} else {
		return json.Marshal(nil)
	}
}

func (p *TemplateId) UnmarshalJSON(b []byte) error {
	var pi *BuildConfigurationShort
	if err := json.Unmarshal(b, &pi); err != nil {
		return err
	}
	if pi == nil {
		*p = ""
	} else {
		*p = TemplateId(pi.ID)
	}
	return nil
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
