package {{ .PackageName }}

// NewDataSource is a helper function to simplify the provider implementation.
func NewDataSource() datasource.DataSource {
	return &DataSource{}
}

// DataSource is the data source implementation for {{ .Name }}
type DataSource struct{
    client   c.Client
    endpoint string
    name     string
}
