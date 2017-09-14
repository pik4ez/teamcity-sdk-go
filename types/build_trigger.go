package types

import (
	"encoding/json"
)

type BuildTrigger struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
}

type BuildTriggers []BuildTrigger

type buildTriggersInput struct {
	Trigger []BuildTrigger `json:"trigger"`
}

func (bt BuildTriggers) MarshalJSON() ([]byte, error) {
	bti := &buildTriggersInput{
		Trigger: bt,
	}
	return json.Marshal(bti)
}

func (bt *BuildTriggers) UnmarshalJSON(b []byte) error {
	var bti buildTriggersInput
	if err := json.Unmarshal(b, &bti); err != nil {
		return err
	}
	if bti.Trigger != nil {
		*bt = bti.Trigger
	} else {
		*bt = make(BuildTriggers, 0)
	}
	return nil
}
