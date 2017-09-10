package gocd

// StagesService exposes calls for interacting with Stage objects in the GoCD API.
type StagesService service

// Stage represents a GoCD Stage object.
type Stage struct {
	Name                  string    `json:"name"`
	FetchMaterials        bool      `json:"fetch_materials"`
	CleanWorkingDirectory bool      `json:"clean_working_directory"`
	NeverCleanupArtifacts bool      `json:"never_cleanup_artifacts"`
	Approval              *Approval `json:"approval,omitempty"`
	EnvironmentVariables  []string  `json:"environment_variables,omitempty"`
	Resources             []string  `json:"resource,omitempty"`
	Jobs                  []*Job    `json:"jobs,omitempty"`
}
