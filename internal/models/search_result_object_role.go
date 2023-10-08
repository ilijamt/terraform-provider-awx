package models

type SearchResultObjectRole struct {
	Count   int64        `json:"count" mapstructure:"count"`
	Results []ObjectRole `json:"results" mapstructure:"results"`
}
