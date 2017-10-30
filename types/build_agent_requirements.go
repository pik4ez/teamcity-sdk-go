package types

import (
	"encoding/json"
)

type BuildAgentRequirement struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
}

type BuildAgentRequirements []BuildAgentRequirement

type buildAgentRequirementsInput struct {
	AgentRequirement []BuildAgentRequirement `json:"agent-requirement"`
}

func (bf BuildAgentRequirements) MarshalJSON() ([]byte, error) {
	bfi := &buildAgentRequirementsInput{
		AgentRequirement: bf,
	}
	return json.Marshal(bfi)
}

func (bf *BuildAgentRequirements) UnmarshalJSON(b []byte) error {
	var bfi buildAgentRequirementsInput
	if err := json.Unmarshal(b, &bfi); err != nil {
		return err
	}
	if bfi.AgentRequirement != nil {
		*bf = bfi.AgentRequirement
	} else {
		*bf = make(BuildAgentRequirements, 0)
	}
	return nil
}
