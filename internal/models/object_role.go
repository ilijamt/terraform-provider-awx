package models

type ObjectRole struct {
	ID          int64  `json:"id" mapstructure:"id"`
	Name        string `json:"name" mapstructure:"name"`
	Description string `json:"description" mapstructure:"description"`
}
