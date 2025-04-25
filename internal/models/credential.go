package models

type Credential struct {
	Inputs         map[string]any `json:"inputs" yaml:"inputs" yaml:"inputs"`
	Name           string         `json:"name" yaml:"name"`
	Description    string         `json:"description" yaml:"description"`
	CredentialType int            `json:"credential_type" yaml:"credential_type"`
	User           int            `json:"user,omitempty" yaml:"user"`
	Team           int            `json:"team,omitempty" yaml:"team"`
	Organization   int            `json:"organization,omitempty" yaml:"organization"`
}
