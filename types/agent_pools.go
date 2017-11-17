package types

import (
	"encoding/json"
)

type AttachedProjects map[string]string

type AgentPools struct {
	ID       int      `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Href     string   `json:"href,omitempty"`
	Projects Projects `json:"projects,omitempty"`
}

type AgentPool struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Href string `json:"href,omitempty"`
}

type AgentPoolAttachment struct {
	ProjectID string `json:"id,omitempty"`
}

type AgentPoolId int

func (v AgentPoolId) MarshalJSON() ([]byte, error) {
	vrs := &AgentPool{
		ID: int(v),
	}
	return json.Marshal(vrs)
}
