package framework_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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
