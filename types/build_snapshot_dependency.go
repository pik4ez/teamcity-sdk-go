package types

import (
	"encoding/json"
)

type BuildSnapshotDependency struct {
	ID              string     `json:"id"`
	Type            string     `json:"type"`
	Properties      Properties `json:"properties"`
	SourceBuildType BuildType  `json:"source-buildType"`
}

type BuildSnapshotDependencies []BuildSnapshotDependency

type buildSnapshotDependenciesInput struct {
	SnapshotDependency []BuildSnapshotDependency `json:"snapshot-dependency"`
}

func (bsd BuildSnapshotDependencies) MarshalJSON() ([]byte, error) {
	bsdi := &buildSnapshotDependenciesInput{
		SnapshotDependency: bsd,
	}
	return json.Marshal(bsdi)
}

func (bsd *BuildSnapshotDependencies) UnmarshalJSON(b []byte) error {
	var bsdi buildSnapshotDependenciesInput
	if err := json.Unmarshal(b, &bsdi); err != nil {
		return err
	}
	if bsdi.SnapshotDependency != nil {
		*bsd = bsdi.SnapshotDependency
	} else {
		*bsd = make(BuildSnapshotDependencies, 0)
	}
	return nil
}
