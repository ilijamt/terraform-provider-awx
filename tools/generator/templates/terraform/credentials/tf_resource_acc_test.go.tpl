package {{ .PackageName }}_test

import (
	"testing"

	r "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/awx/credential/{{ .PackageName }}"
	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/provider"
	"github.com/ilijamt/terraform-provider-awx/version"
)

func TestAccResource(t *testing.T) {
	client, _ := c.NewTestingClient(t)

	resource.Test(
		t,
		resource.TestCase{
			PreCheck: func() { provider.TestAccPreCheck(t) },
			ProtoV6ProviderFactories: provider.TestFactories(
                t, "awx", nil, client, version.Version,
                []func() r.Resource{ {{ .PackageName }}.NewResource },
                provider.TestEmptyDataSources(t),
            ),
			Steps: []resource.TestStep{},
		},
	)
}
