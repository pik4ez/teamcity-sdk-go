package types

import (
	"encoding/json"
	"fmt"
)

// Build represents a TeamCity build, along with its metadata.
type Build struct {
	ID          int64
	BuildTypeID string
	BuildType struct {
		ID          string
		Name        string
		Description string
		ProjectName string
		ProjectID   string
		HREF        string
		WebURL      string
	}
	Triggered struct {
		Type string
		Date JSONTime
		User struct {
			Username string
		}
	}
	Changes struct {
		Change []Change
	}

	QueuedDate    JSONTime
	StartDate     JSONTime
	FinishDate    JSONTime
	Number        string
	Status        string
	StatusText    string
	State         string
	BranchName    string
	Personal      bool
	Running       bool
	Pinned        bool
	DefaultBranch bool
	HREF          string
	WebURL        string
	Agent struct {
		ID     int64
		Name   string
		TypeID int64
		HREF   string
	}

	ProblemOccurrences struct {
		ProblemOccurrence []ProblemOccurrence
	}

	TestOccurrences struct {
		TestOccurrence []TestOccurrence
	}

	Tags tagsInput `json:"tags,omitempty"`

	RunningInfo RunningInfo `json:"running-info,omitempty"`

	Properties Properties `json:"properties"`

	ArtifactDependencies Dependencies `json:"artifact-dependencies,omitempty"`
	SnapshotDependencies Dependencies `json:"snapshot-dependencies,omitempty"`
}

type Tags []string

type tagInput struct {
	Name string `json:"name"`
}

type tagsInput struct {
	Tag []tagInput `json:"tag"`
}

func (tags Tags) MarshalJSON() ([]byte, error) {
	tagInputs := make([]tagInput, len(tags))
	for idx, tag := range tags {
		tagInputs[idx] = tagInput{tag}
	}
	ti := &tagsInput{
		Tag: tagInputs,
	}
	return json.Marshal(ti)
}

func (tags *Tags) UnmarshalJSON(b []byte) error {
	var ti tagsInput
	if err := json.Unmarshal(b, &ti); err != nil {
		return err
	}
	if ti.Tag != nil {
		*tags = make(Tags, len(ti.Tag))
		for idx, tag := range ti.Tag {
			(*tags)[idx] = tag.Name
		}
	} else {
		*tags = make(Tags, 0)
	}
	return nil
}

type RunningInfo struct {
	PercentageComplete    int    `json:"percentageComplete"`
	ElapsedSeconds        int    `json:"elapsedSeconds"`
	EstimatedTotalSeconds int    `json:"estimatedTotalSeconds"`
	CurrentStageText      string `json:"currentStageText"`
	Outdated              bool   `json:"outdated"`
	ProbablyHanging       bool   `json:"probablyHanging"`
}

type Dependencies []Dependency

type Dependency struct {
	ID          int    `json:"id"`
	BuildTypeID string `json:"buildTypeID"`
	Number      string `json:"number"`
	Status      string `json:"status"`
	State       string `json:"state"`
	BranchName  string `json:"branchName"`
	HREF        string `json:"href"`
	WebURL      string `json:"webUrl"`
}

func (dependencies *Dependencies) UnmarshalJSON(b []byte) error {
	var depInput struct {
		Count int          `json:"count,omitempty"`
		Build []Dependency `json:"build,omitempty"`
	}

	if err := json.Unmarshal(b, &depInput); err != nil {
		return err
	}

	if depInput.Count != 0 {
		*dependencies = Dependencies(depInput.Build)
	} else {
		*dependencies = make(Dependencies, 0)
	}

	return nil
}

func (b *Build) String() string {
	return fmt.Sprintf("Build %d, %#v state=%s", b.ID, b.ComputedState(), b.State)
}

type State int

const (
	Unknown  = State(iota)
	Queued
	Started
	Finished
)

func (b *Build) ComputedState() State {
	if b.QueuedDate == "" {
		return Unknown
	}
	if b.StartDate == "" {
		return Queued
	}
	if b.FinishDate == "" {
		return Started
	}
	return Finished
}
