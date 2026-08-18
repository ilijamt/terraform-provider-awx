package framework_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

func TestAssociateDisassociateResource_Schema(t *testing.T) {
	tests := []struct {
		name            string
		cfg             framework.AssociateDisassociateConfig
		wantParentDesc  string
		wantChildDesc   string
		wantHasOption   bool
		wantDeprecation string
	}{
		{
			name: "host group default",
			cfg: framework.AssociateDisassociateConfig{
				TypeName: "host_associate_group", Endpoint: "/api/v2/hosts/%d/groups/",
				ParentName: "Host", ParentIDAttr: "host_id",
				ChildName: "Group", ChildIDAttr: "group_id",
				Deprecated: true,
			},
			wantParentDesc:  "Database ID for this Host.",
			wantChildDesc:   "Database ID of the group to assign.",
			wantHasOption:   false,
			wantDeprecation: "This resource has been deprecated and will be removed in a future release.",
		},
		{
			name: "job template notification",
			cfg: framework.AssociateDisassociateConfig{
				TypeName: "job_template_associate_notification_template", Endpoint: "/api/v2/job_templates/%d/notification_templates_%s/",
				ParentName: "JobTemplate", ParentIDAttr: "job_template_id",
				ChildName: "NotificationTemplate", ChildIDAttr: "notification_template_id",
				AssociateType: "notification_job_template", Deprecated: true,
			},
			wantParentDesc:  "Database ID for this JobTemplate.",
			wantChildDesc:   "Database ID of the notificationtemplate to assign.",
			wantHasOption:   true,
			wantDeprecation: "This resource has been deprecated and will be removed in a future release.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := framework.NewAssociateDisassociateResource(tt.cfg)
			schemaProvider, ok := r.(interface {
				Schema(context.Context, resource.SchemaRequest, *resource.SchemaResponse)
			})
			require.True(t, ok)

			resp := &resource.SchemaResponse{}
			schemaProvider.Schema(context.Background(), resource.SchemaRequest{}, resp)

			assert.Equal(t, tt.wantDeprecation, resp.Schema.DeprecationMessage)

			parent, ok := resp.Schema.Attributes[tt.cfg.ParentIDAttr]
			require.True(t, ok, "missing parent ID attribute")
			assert.Equal(t, tt.wantParentDesc, parent.GetDescription())
			assert.True(t, parent.IsRequired())

			child, ok := resp.Schema.Attributes[tt.cfg.ChildIDAttr]
			require.True(t, ok, "missing child ID attribute")
			assert.Equal(t, tt.wantChildDesc, child.GetDescription())
			assert.True(t, child.IsRequired())

			_, hasOption := resp.Schema.Attributes["option"]
			assert.Equal(t, tt.wantHasOption, hasOption)
		})
	}
}

func TestAssociateDisassociateResource_ConfigValidators(t *testing.T) {
	tests := []struct {
		name    string
		cfg     framework.AssociateDisassociateConfig
		wantLen int
	}{
		{
			name: "default has one validator",
			cfg: framework.AssociateDisassociateConfig{
				ParentIDAttr: "host_id", ChildIDAttr: "group_id",
			},
			wantLen: 1,
		},
		{
			name: "notification has one validator (covers all three attrs)",
			cfg: framework.AssociateDisassociateConfig{
				ParentIDAttr: "job_template_id", ChildIDAttr: "notification_template_id",
				AssociateType: "notification_job_template",
			},
			wantLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := framework.NewAssociateDisassociateResource(tt.cfg)
			cv, ok := r.(interface {
				ConfigValidators(context.Context) []resource.ConfigValidator
			})
			require.True(t, ok)
			assert.Len(t, cv.ConfigValidators(context.Background()), tt.wantLen)
		})
	}
}

func TestAssociateDisassociateResource_Metadata(t *testing.T) {
	r := framework.NewAssociateDisassociateResource(framework.AssociateDisassociateConfig{
		TypeName: "host_associate_group",
	})
	mdProvider, ok := r.(interface {
		Metadata(context.Context, resource.MetadataRequest, *resource.MetadataResponse)
	})
	require.True(t, ok)

	resp := &resource.MetadataResponse{}
	mdProvider.Metadata(context.Background(), resource.MetadataRequest{ProviderTypeName: "awx"}, resp)
	assert.Equal(t, "awx_host_associate_group", resp.TypeName)
}

func readTestResource(t *testing.T, req *mockRequester) (*framework.AssociateDisassociateResource, tfsdk.State) {
	t.Helper()
	r, ok := framework.NewAssociateDisassociateResource(framework.AssociateDisassociateConfig{
		TypeName:   "workflow_job_template_node_associate_success_node",
		Endpoint:   "/api/v2/workflow_job_template_nodes/%d/success_nodes/",
		ParentName: "WorkflowJobTemplateNode", ParentIDAttr: "workflow_job_template_node_id",
		ChildName: "SuccessNode", ChildIDAttr: "success_node_id",
	}).(*framework.AssociateDisassociateResource)
	require.True(t, ok)
	r.Client = req

	schemaResp := &resource.SchemaResponse{}
	r.Schema(context.Background(), resource.SchemaRequest{}, schemaResp)

	return r, tfsdk.State{
		Schema: schemaResp.Schema,
		Raw: tftypes.NewValue(schemaResp.Schema.Type().TerraformType(context.Background()), map[string]tftypes.Value{
			"workflow_job_template_node_id": tftypes.NewValue(tftypes.Number, 7),
			"success_node_id":               tftypes.NewValue(tftypes.Number, 9),
		}),
	}
}

func TestAssociateDisassociateResource_Read(t *testing.T) {
	notFound := &mockRequester{
		newRequestFunc: func(context.Context, string, string, io.Reader) (*http.Request, error) {
			return &http.Request{}, nil
		},
		doFunc: func(context.Context, *http.Request) (map[string]any, error) {
			return nil, &c.StatusError{StatusCode: http.StatusNotFound, URI: "/api/v2/x/", Body: "gone"}
		},
	}

	for _, tt := range []struct {
		name        string
		requester   *mockRequester
		wantRemoved bool
		wantError   bool
	}{
		{name: "still associated keeps state", requester: successRequester(map[string]any{"count": json.Number("1")})},
		{name: "unlinked in AWX drops state", requester: successRequester(map[string]any{"count": json.Number("0")}), wantRemoved: true},
		{name: "missing parent drops state", requester: notFound, wantRemoved: true},
		{name: "other failures surface", requester: failDo(), wantError: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r, state := readTestResource(t, tt.requester)
			resp := &resource.ReadResponse{State: state}
			r.Read(context.Background(), resource.ReadRequest{State: state}, resp)

			assert.Equal(t, tt.wantError, resp.Diagnostics.HasError())
			assert.Equal(t, tt.wantRemoved, resp.State.Raw.IsNull())
		})
	}
}

func TestAssociateDisassociateResource_ReadFiltersByChildID(t *testing.T) {
	var got string
	req := &mockRequester{
		newRequestFunc: func(_ context.Context, _, endpoint string, _ io.Reader) (*http.Request, error) {
			got = endpoint
			return &http.Request{}, nil
		},
		doFunc: func(context.Context, *http.Request) (map[string]any, error) {
			return map[string]any{"count": json.Number("1")}, nil
		},
	}
	r, state := readTestResource(t, req)
	r.Read(context.Background(), resource.ReadRequest{State: state}, &resource.ReadResponse{State: state})

	assert.Equal(t, "/api/v2/workflow_job_template_nodes/7/success_nodes/?id=9", got)
}
