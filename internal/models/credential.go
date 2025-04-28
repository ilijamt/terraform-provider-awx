package models

type Credential struct {
	Inputs         map[string]any `json:"inputs" yaml:"inputs" yaml:"inputs"`
	Name           string         `json:"name" yaml:"name"`
	Description    string         `json:"description" yaml:"description"`
	CredentialType int64          `json:"credential_type" yaml:"credential_type"`
	User           int64          `json:"user,omitzero" yaml:"user"`
	Organization   *int64         `json:"organization,omitzero" yaml:"organization"`
}
