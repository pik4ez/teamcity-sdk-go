package types

import (
	"encoding/json"
)

type BuildFeature struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
}

type BuildFeatures []BuildFeature

type buildFeaturesInput struct {
	Feature []BuildFeature `json:"feature"`
}

func (bf BuildFeatures) MarshalJSON() ([]byte, error) {
	bfi := &buildFeaturesInput{
		Feature: bf,
	}
	return json.Marshal(bfi)
}

func (bf *BuildFeatures) UnmarshalJSON(b []byte) error {
	var bfi buildFeaturesInput
	if err := json.Unmarshal(b, &bfi); err != nil {
		return err
	}
	if bfi.Feature != nil {
		*bf = bfi.Feature
	} else {
		*bf = make(BuildFeatures, 0)
	}
	return nil
}
