package types

import (
	"encoding/json"
)

type VcsRoot struct {
	ID         string     `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	VcsName    string     `json:"vcsName,omitempty"`
	Href       string     `json:"href,omitempty"`
	ProjectID  ProjectId  `json:"project,omitempty"`
	Properties Properties `json:"properties,omitempty"`
}

type VcsRootShort struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Href string `json:"href"`
}

type VcsRootId string

func (v VcsRootId) MarshalJSON() ([]byte, error) {
	vrs := &VcsRootShort{
		ID: string(v),
	}
	return json.Marshal(vrs)
}

func (v *VcsRootId) UnmarshalJSON(b []byte) error {
	var vrs VcsRootShort
	if err := json.Unmarshal(b, &vrs); err != nil {
		return err
	}
	*v = VcsRootId(vrs.ID)
	return nil
}

type VcsRoots []VcsRootShort

type vcsRootsInput struct {
	VcsRoot []VcsRootShort `json:"vcs-root"`
}

func (vc VcsRoots) MarshalJSON() ([]byte, error) {
	vci := &vcsRootsInput{
		VcsRoot: vc,
	}
	return json.Marshal(vci)
}

func (vc *VcsRoots) UnmarshalJSON(b []byte) error {
	var vci vcsRootsInput
	if err := json.Unmarshal(b, &vci); err != nil {
		return err
	}
	if vci.VcsRoot != nil {
		*vc = vci.VcsRoot
	} else {
		*vc = make(VcsRoots, 0)
	}
	return nil
}
