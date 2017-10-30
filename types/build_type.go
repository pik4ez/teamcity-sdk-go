package types

type BuildType struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ProjectName string `json:"projectName"`
	ProjectID   string `json:"projectId"`
	Href        string `json:"href,omitempty"`
	webURL      string `json:"webUrl,omitempty"`
}
