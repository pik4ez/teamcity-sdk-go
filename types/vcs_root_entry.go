package types

import (
	"encoding/json"
)

type VcsRootEntry struct {
	ID            string    `json:"id,omitempty"`
	VcsRootID     VcsRootId `json:"vcs-root"`
	CheckoutRules string    `json:"checkout-rules"`
}

type VcsRootEntries []VcsRootEntry

type vcsRootEntriesInput struct {
	VcsRootEntry []VcsRootEntry `json:"vcs-root-entry"`
}

func (vre VcsRootEntries) MarshalJSON() ([]byte, error) {
	vrei := &vcsRootEntriesInput{
		VcsRootEntry: vre,
	}
	return json.Marshal(vrei)
}

func (vre *VcsRootEntries) UnmarshalJSON(b []byte) error {
	var vrei vcsRootEntriesInput
	if err := json.Unmarshal(b, &vrei); err != nil {
		return err
	}
	if vrei.VcsRootEntry != nil {
		*vre = vrei.VcsRootEntry
	} else {
		*vre = make(VcsRootEntries, 0)
	}
	return nil
}
