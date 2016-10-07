package types

import (
	"encoding/json"
)

type Project struct {
	ID                  string              `json:"id,omitempty"`
	Name                string              `json:"name"`
	Description         string              `json:"description,omitempty"`
	Href                string              `json:"href,omitempty"`
	WebUrl              string              `json:"webUrl,omitempty"`
	ParentProjectID     ProjectId           `json:"parentProject,omitempty"`
	BuildConfigurations BuildConfigurations `json:"buildTypes,omitempty"`
	Templates           BuildConfigurations `json:"templates,omitempty"`
	Parameters          Parameters          `json:"parameters,omitempty"`
	Projects            Projects            `json:"projects,omitempty"`
}

type ProjectId string

type ProjectShort struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description,omitempty"`
	ParentProjectID string `json:"parentProjectId"`
	Href            string `json:"href"`
	WebURL          string `json:"webUrl"`
}

func (p ProjectId) MarshalJSON() ([]byte, error) {
	pi := &ProjectShort{
		ID: string(p),
	}
	return json.Marshal(pi)
}

func (p *ProjectId) UnmarshalJSON(b []byte) error {
	var pi ProjectShort
	if err := json.Unmarshal(b, &pi); err != nil {
		return err
	}
	*p = ProjectId(pi.ID)
	return nil
}

type Projects map[string]Project

type projectsInput struct {
	Project []Project `json:"project"`
}

func (p Projects) MarshalJSON() ([]byte, error) {
	pi := &projectsInput{
		Project: make([]Project, 0),
	}
	for _, value := range p {
		pi.Project = append(pi.Project, value)
	}
	return json.Marshal(pi)
}

func (p *Projects) UnmarshalJSON(b []byte) error {
	var pi projectsInput
	if err := json.Unmarshal(b, &pi); err != nil {
		return err
	}
	m := make(Projects)
	for _, proj := range pi.Project {
		m[proj.ID] = proj
	}
	*p = m
	return nil
}
