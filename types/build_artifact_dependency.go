package types

import (
	"encoding/json"
)

type BuildArtifactDependency struct {
	ID              string     `json:"id"`
	Type            string     `json:"type"`
	Properties      Properties `json:"properties"`
	SourceBuildType BuildType  `json:"source-buildType"`
}

type BuildArtifactDependencies []BuildArtifactDependency

type buildArtifactDependenciesInput struct {
	ArtifactDependency []BuildArtifactDependency `json:"artifact-dependency"`
}

func (bsd BuildArtifactDependencies) MarshalJSON() ([]byte, error) {
	bsdi := &buildArtifactDependenciesInput{
		ArtifactDependency: bsd,
	}
	return json.Marshal(bsdi)
}

func (bsd *BuildArtifactDependencies) UnmarshalJSON(b []byte) error {
	var bsdi buildArtifactDependenciesInput
	if err := json.Unmarshal(b, &bsdi); err != nil {
		return err
	}
	if bsdi.ArtifactDependency != nil {
		*bsd = bsdi.ArtifactDependency
	} else {
		*bsd = make(BuildArtifactDependencies, 0)
	}
	return nil
}
