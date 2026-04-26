package framework_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

// modelStub satisfies framework.ResourceModel[modelStub, bodyStub] for the
// generic resource type parameters. The CRUD methods aren't exercised in
// these tests — we only need Configure to wire up.
type modelStub struct{}

func (m *modelStub) Clone() modelStub                                                   { return *m }
func (m *modelStub) BodyRequest() *bodyStub                                             { return &bodyStub{} }
func (m *modelStub) UpdateFromApiData(map[string]any) (diag.Diagnostics, error)         { return nil, nil }

type bodyStub struct{}

func TestGenericResource_Configure_RunsOnConfigure(t *testing.T) {
	tests := []struct {
		name           string
		onConfigure    framework.ConfigureFunc
		providerData   any
		wantInvoked    bool
		wantDiagsError bool
	}{
		{
			name: "OnConfigure runs and stashes value",
			onConfigure: func(ctx context.Context, c framework.Requester) diag.Diagnostics {
				assert.NotNil(t, c, "client must be non-nil when OnConfigure fires")
				return nil
			},
			providerData:   successRequester(nil),
			wantInvoked:    true,
			wantDiagsError: false,
		},
		{
			name: "OnConfigure error surfaces as diagnostic",
			onConfigure: func(_ context.Context, _ framework.Requester) diag.Diagnostics {
				d := diag.Diagnostics{}
				d.AddError("lookup failed", "couldn't find namespace")
				return d
			},
			providerData:   successRequester(nil),
			wantInvoked:    true,
			wantDiagsError: true,
		},
		{
			name: "OnConfigure skipped when client is nil",
			onConfigure: func(_ context.Context, _ framework.Requester) diag.Diagnostics {
				t.Fatal("OnConfigure must not run when ProviderData is nil")
				return nil
			},
			providerData:   nil,
			wantInvoked:    false,
			wantDiagsError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invoked := false
			r := &framework.GenericResource[modelStub, bodyStub, *modelStub]{
				ResourceBase: framework.ResourceBase{
					ProviderBase: framework.ProviderBase{TypeName: "stub", Endpoint: "/api/v2/stub/"},
				},
				Cfg: framework.ResourceCfg[modelStub, bodyStub]{
					Schema: rschema.Schema{},
					OnConfigure: func(ctx context.Context, c framework.Requester) diag.Diagnostics {
						invoked = true
						return tt.onConfigure(ctx, c)
					},
				},
			}

			req := resource.ConfigureRequest{ProviderData: tt.providerData}
			resp := &resource.ConfigureResponse{}
			r.Configure(context.Background(), req, resp)

			assert.Equal(t, tt.wantInvoked, invoked)
			if tt.wantDiagsError {
				require.True(t, resp.Diagnostics.HasError())
			} else {
				assert.False(t, resp.Diagnostics.HasError())
			}
		})
	}
}
