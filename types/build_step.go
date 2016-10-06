package types

import (
	"encoding/json"
)

type BuildStep struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Name       string     `json:"name"`
	Properties Properties `json:"properties"`
}

type BuildSteps []BuildStep

type buildStepsInput struct {
	Step []BuildStep `json:"step"`
}

func (bs BuildSteps) MarshalJSON() ([]byte, error) {
	bsi := &buildStepsInput{
		Step: bs,
	}
	return json.Marshal(bsi)
}

func (bs *BuildSteps) UnmarshalJSON(b []byte) error {
	var bsi buildStepsInput
	if err := json.Unmarshal(b, &bsi); err != nil {
		return err
	}
	if bsi.Step != nil {
		*bs = bsi.Step
	} else {
		*bs = make(BuildSteps, 0)
	}
	return nil
}
