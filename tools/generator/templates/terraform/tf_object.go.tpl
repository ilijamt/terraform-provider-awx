{{- /*
tf_object.go.tpl is the entry template for a single regular (non-credential)
generated AWX object. It emits the package declaration, a kitchen-sink import
block (goimports trims unused entries after generation), and stitches together
three partials:

  - model_section        — typed model + body request struct + reader
  - resource_section     — schema + GenericResource binding (gated by NoTerraformResource)
  - data_source_section  — schema + GenericDataSource binding (gated by NoTerraformDataSource)

The data source section uses the `dschema` alias because the resource section
binds `schema` to resource/schema.
*/ -}}
package {{ .PackageName }}

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

{{ template "model_section" . }}

{{ if not .NoTerraformResource }}
{{ template "resource_section" . }}
{{ end }}

{{ if not .NoTerraformDataSource }}
{{ template "data_source_section" . }}
{{ end }}
