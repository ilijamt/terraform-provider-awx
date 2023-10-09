package internal

type SearchResults struct {
	Count    int              `json:"count" mapstructure:"count"`
	Next     string           `json:"next" mapstructure:"next"`
	Previous string           `json:"previous" mapstructure:"previous"`
	Results  []map[string]any `json:"results" mapstructure:"results"`
}
