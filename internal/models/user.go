package models

type User struct {
	ID              int    `json:"id"`
	IsSuperuser     bool   `json:"is_superuser"`
	IsSystemAuditor bool   `json:"is_system_auditor"`
	Username        string `json:"username"`
	EMail           string `json:"email"`
}
