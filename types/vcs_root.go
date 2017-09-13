package types

import (
	"encoding/json"
)

type VcsRoot struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	VcsName    string     `json:"vcsName"`
	Href       string     `json:"href"`
	ProjectID  string     `json:"projectId"`
	Properties Properties `json:"properties"`
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
