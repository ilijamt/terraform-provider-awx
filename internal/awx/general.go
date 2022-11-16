package awx

type objectRole struct {
	ID          int64  `json:"id" mapstructure:"id"`
	Name        string `json:"name" mapstructure:"name"`
	Description string `json:"description" mapstructure:"description"`
}

type searchResultObjectRole struct {
	Count   int64        `json:"count" mapstructure:"count"`
	Results []objectRole `json:"results" mapstructure:"results"`
}
