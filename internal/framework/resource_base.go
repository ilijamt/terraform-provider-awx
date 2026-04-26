package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// ResourceBase provides the shared fields and methods for generated resources.
// Configure and Metadata are promoted and match the Terraform resource interfaces.
type ResourceBase struct {
	ProviderBase
}

func (b *ResourceBase) Configure(_ context.Context, request resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	b.configureClient(request.ProviderData)
}

func (b *ResourceBase) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_" + b.TypeName
}
