package framework

import (
	"context"
	"sync/atomic"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// CredentialTypeLookup pairs an atomic ID slot with an OnConfigure that
// populates it. One value is shared between a typed-credential resource and
// its matching data source so a single namespace lookup at Configure time
// covers both. Use NewCredentialTypeLookup to construct.
type CredentialTypeLookup struct {
	id atomic.Int64
}

// NewCredentialTypeLookup returns a *CredentialTypeLookup. The returned value
// must be stored in a package-level var so its OnConfigure closure and Load
// observers see the same atomic.
func NewCredentialTypeLookup() *CredentialTypeLookup { return &CredentialTypeLookup{} }

// OnConfigure returns a ConfigureFunc that resolves the AWX credential_type
// ID for namespace and stores it on the lookup.
func (l *CredentialTypeLookup) OnConfigure(namespace string) ConfigureFunc {
	return func(ctx context.Context, client Requester) diag.Diagnostics {
		id, d := LookupCredentialTypeIDByNamespace(ctx, client, namespace)
		if d.HasError() {
			return d
		}
		l.id.Store(id)
		return d
	}
}

// Load returns the resolved credential_type ID, or 0 if Configure hasn't run.
func (l *CredentialTypeLookup) Load() int64 { return l.id.Load() }

// CredentialBaseResourceAttrs returns the nine ownership/identity attributes
// shared by every typed-credential resource: id, name, description,
// organization, team, user, kind, managed, credential_type. Returns a fresh
// map per call so callers can layer per-credential-type attrs on top.
func CredentialBaseResourceAttrs() map[string]schema.Attribute {
	ownerValidator := []validator.Int64{
		int64validator.ExactlyOneOf(
			path.MatchRoot("organization"),
			path.MatchRoot("team"),
			path.MatchRoot("user"),
		),
	}
	owner := func(desc string) schema.Int64Attribute {
		return schema.Int64Attribute{
			Description: desc,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
			Validators: ownerValidator,
		}
	}
	return map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Description: "Database ID of this credential.",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Description: "Name of this credential.",
			Required:    true,
			Validators: []validator.String{
				stringvalidator.LengthAtMost(512),
			},
		},
		"description": schema.StringAttribute{
			Description: "Optional description of this credential.",
			Optional:    true,
			Computed:    true,
			Default:     stringdefault.StaticString(""),
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"organization": owner("Inherit permissions from organization roles. Mutually exclusive with team and user."),
		"team":         owner("Write-only field used to add team to owner role. Mutually exclusive with organization and user. Only valid for creation."),
		"user":         owner("Write-only field used to add user to owner role. Mutually exclusive with organization and team. Only valid for creation."),
		"credential_type": schema.Int64Attribute{
			Description: "Resolved AWX credential_type ID for this credential's namespace. Computed at Configure time.",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"kind": schema.StringAttribute{
			Description: "AWX credential kind — the namespace of the credential type (e.g. aws / ssh / vault).",
			Computed:    true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"managed": schema.BoolAttribute{
			Description: "Whether AWX considers this a managed credential.",
			Computed:    true,
		},
	}
}

// CredentialBaseDataSourceAttrs is the data-source variant of
// CredentialBaseResourceAttrs. id and name are Optional+Computed with an
// ExactlyOneOf validator; the rest are Computed-only.
func CredentialBaseDataSourceAttrs() map[string]dschema.Attribute {
	idOrName := []path.Expression{path.MatchRoot("id"), path.MatchRoot("name")}
	return map[string]dschema.Attribute{
		"id": dschema.Int64Attribute{
			Description: "Database ID of this credential.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Int64{
				int64validator.ExactlyOneOf(idOrName...),
			},
		},
		"name": dschema.StringAttribute{
			Description: "Name of this credential.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(idOrName...),
			},
		},
		"description":     dschema.StringAttribute{Description: "Optional description of this credential.", Computed: true},
		"organization":    dschema.Int64Attribute{Description: "Owning organization ID.", Computed: true},
		"team":            dschema.Int64Attribute{Description: "Owning team ID (write-only on create).", Computed: true},
		"user":            dschema.Int64Attribute{Description: "Owning user ID (write-only on create).", Computed: true},
		"credential_type": dschema.Int64Attribute{Description: "Resolved AWX credential_type ID for this credential's namespace.", Computed: true},
		"kind":            dschema.StringAttribute{Description: "AWX credential kind — the namespace of the credential type (e.g. aws / ssh / vault).", Computed: true},
		"managed":         dschema.BoolAttribute{Description: "Whether AWX considers this a managed credential.", Computed: true},
	}
}
