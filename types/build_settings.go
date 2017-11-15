package types

import (
	"encoding/json"
)

type BuildSetting struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type BuildSettings []BuildSetting

type buildSettingsInput struct {
	Setting []BuildSetting `json:"property"`
}

func (s BuildSettings) MarshalJSON() ([]byte, error) {
	si := &buildSettingsInput{
		Setting: s,
	}
	return json.Marshal(si)
}

func (s *BuildSettings) UnmarshalJSON(b []byte) error {
	var si buildSettingsInput
	if err := json.Unmarshal(b, &si); err != nil {
		return err
	}
	if si.Setting != nil {
		*s = si.Setting
	} else {
		*s = make(BuildSettings, 0)
	}
	return nil
}
