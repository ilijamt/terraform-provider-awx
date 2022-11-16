package {{ .PackageName }}

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	p "path"
	"strconv"
	"strings"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/mitchellh/mapstructure"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

{{ block "tf_model" . }}{{ end }}

{{ if not $.Config.NoTerraformDataSource }}
{{ block "tf_data_source" . }}{{ end }}
{{ end }}

{{ if not $.Config.NoTerraformResource }}
{{ block "tf_resource" . }}{{ end }}
{{ end }}

{{ if $.Config.HasObjectRoles }}
{{ block "tf_resource_object_role" . }}{{ end }}
{{ end }}

{{ range $key := $.Config.AssociateDisassociateGroups }}
{{ block "tf_associate_disassociate" . }}{{ end }}
{{ end }}

{{ if $.Config.HasSurveySpec }}
{{ block "tf_survey_spec" . }}{{ end }}
{{ end }}