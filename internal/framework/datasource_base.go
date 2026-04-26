package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// DataSourceBase provides the shared fields and methods for generated data sources.
// Configure and Metadata are promoted and match the Terraform datasource interfaces.
type DataSourceBase struct {
	ProviderBase
	Validators []datasource.ConfigValidator
}

func (b *DataSourceBase) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	b.configureClient(req.ProviderData)
}

func (b *DataSourceBase) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + b.TypeName
}

func (b *DataSourceBase) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	if b.Validators != nil {
		return b.Validators
	}
	return []datasource.ConfigValidator{}
}
