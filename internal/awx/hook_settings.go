package awx

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func hookSettingsAuthAzureADOauth2(ctx context.Context, source Source, callee Callee, orig *settingsAuthAzureADOauth2TerraformModel, state *settingsAuthAzureADOauth2TerraformModel) (err error) {
	if source == SourceResource && (state == nil || orig == nil) && (callee == CalleeUpdate || callee == CalleeCreate || callee == CalleeRead) {
		return fmt.Errorf("state and orig required for resource")
	}

	if source == SourceResource && callee == CalleeCreate ||
		(state.SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET.Equal(types.StringValue("$encrypted$")) &&
			(source == SourceResource && (callee == CalleeUpdate || callee == CalleeRead)) &&
			!orig.SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET.IsNull()) {
		state.SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET = orig.SOCIAL_AUTH_AZUREAD_OAUTH2_SECRET
	}

	return nil
}

func hookSettingsAuthGithub(ctx context.Context, source Source, callee Callee, orig *settingsAuthGithubTerraformModel, state *settingsAuthGithubTerraformModel) (err error) {
	if source == SourceResource && (state == nil || orig == nil) && (callee == CalleeUpdate || callee == CalleeCreate || callee == CalleeRead) {
		return fmt.Errorf("state and orig required for resource")
	}

	if source == SourceResource && callee == CalleeCreate ||
		(state.SOCIAL_AUTH_GITHUB_SECRET.Equal(types.StringValue("$encrypted$")) &&
			(source == SourceResource && (callee == CalleeUpdate || callee == CalleeRead)) &&
			!orig.SOCIAL_AUTH_GITHUB_SECRET.IsNull()) {
		state.SOCIAL_AUTH_GITHUB_SECRET = orig.SOCIAL_AUTH_GITHUB_SECRET
	}

	return nil
}

func hookSettingsAuthGithubEnterprise(ctx context.Context, source Source, callee Callee, orig *settingsAuthGithubEnterpriseTerraformModel, state *settingsAuthGithubEnterpriseTerraformModel) (err error) {
	if source == SourceResource && (state == nil || orig == nil) && (callee == CalleeUpdate || callee == CalleeCreate || callee == CalleeRead) {
		return fmt.Errorf("state and orig required for resource")
	}

	if source == SourceResource && callee == CalleeCreate ||
		(state.SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET.Equal(types.StringValue("$encrypted$")) &&
			(source == SourceResource && (callee == CalleeUpdate || callee == CalleeRead)) &&
			!orig.SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET.IsNull()) {
		state.SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET = orig.SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET
	}

	return nil
}

func hookSettingsAuthGithubEnterpriseOrg(ctx context.Context, source Source, callee Callee, orig *settingsAuthGithubEnterpriseOrgTerraformModel, state *settingsAuthGithubEnterpriseOrgTerraformModel) (err error) {
	if source == SourceResource && (state == nil || orig == nil) && (callee == CalleeUpdate || callee == CalleeCreate || callee == CalleeRead) {
		return fmt.Errorf("state and orig required for resource")
	}

	if source == SourceResource && callee == CalleeCreate ||
		(state.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET.Equal(types.StringValue("$encrypted$")) &&
			(source == SourceResource && (callee == CalleeUpdate || callee == CalleeRead)) &&
			!orig.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET.IsNull()) {
		state.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET = orig.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET
	}

	return nil
}

func hookSettingsAuthGithubEnterpriseTeam(ctx context.Context, source Source, callee Callee, orig *settingsAuthGithubEnterpriseTeamTerraformModel, state *settingsAuthGithubEnterpriseTeamTerraformModel) (err error) {
	if source == SourceResource && (state == nil || orig == nil) && (callee == CalleeUpdate || callee == CalleeCreate || callee == CalleeRead) {
		return fmt.Errorf("state and orig required for resource")
	}

	if source == SourceResource && callee == CalleeCreate ||
		(state.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET.Equal(types.StringValue("$encrypted$")) &&
			(source == SourceResource && (callee == CalleeUpdate || callee == CalleeRead)) &&
			!orig.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET.IsNull()) {
		state.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET = orig.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET
	}

	return nil
}

func hookSettingsAuthGithubOrg(ctx context.Context, source Source, callee Callee, orig *settingsAuthGithubOrgTerraformModel, state *settingsAuthGithubOrgTerraformModel) (err error) {
	if source == SourceResource && (state == nil || orig == nil) && (callee == CalleeUpdate || callee == CalleeCreate || callee == CalleeRead) {
		return fmt.Errorf("state and orig required for resource")
	}

	if source == SourceResource && callee == CalleeCreate ||
		(state.SOCIAL_AUTH_GITHUB_ORG_SECRET.Equal(types.StringValue("$encrypted$")) &&
			(source == SourceResource && (callee == CalleeUpdate || callee == CalleeRead)) &&
			!orig.SOCIAL_AUTH_GITHUB_ORG_SECRET.IsNull()) {
		state.SOCIAL_AUTH_GITHUB_ORG_SECRET = orig.SOCIAL_AUTH_GITHUB_ORG_SECRET
	}

	return nil
}

func hookSettingsAuthGithubTeam(ctx context.Context, source Source, callee Callee, orig *settingsAuthGithubTeamTerraformModel, state *settingsAuthGithubTeamTerraformModel) (err error) {
	if source == SourceResource && (state == nil || orig == nil) && (callee == CalleeUpdate || callee == CalleeCreate || callee == CalleeRead) {
		return fmt.Errorf("state and orig required for resource")
	}

	if source == SourceResource && callee == CalleeCreate ||
		(state.SOCIAL_AUTH_GITHUB_TEAM_SECRET.Equal(types.StringValue("$encrypted$")) &&
			(source == SourceResource && (callee == CalleeUpdate || callee == CalleeRead)) &&
			!orig.SOCIAL_AUTH_GITHUB_TEAM_SECRET.IsNull()) {
		state.SOCIAL_AUTH_GITHUB_TEAM_SECRET = orig.SOCIAL_AUTH_GITHUB_TEAM_SECRET
	}

	return nil
}

func hookSettingsSaml(ctx context.Context, source Source, callee Callee, orig *settingsAuthSAMLTerraformModel, state *settingsAuthSAMLTerraformModel) (err error) {
	if source == SourceResource && (state == nil || orig == nil) && (callee == CalleeUpdate || callee == CalleeCreate || callee == CalleeRead) {
		return fmt.Errorf("state and orig required for resource")
	}

	state.SOCIAL_AUTH_SAML_SP_PUBLIC_CERT = types.StringValue(state.SOCIAL_AUTH_SAML_SP_PUBLIC_CERT.ValueString() + "\n")
	if source == SourceResource && callee == CalleeCreate ||
		(state.SOCIAL_AUTH_SAML_SP_PRIVATE_KEY.Equal(types.StringValue("$encrypted$")) &&
			(source == SourceResource && (callee == CalleeUpdate || callee == CalleeRead)) &&
			!orig.SOCIAL_AUTH_SAML_SP_PRIVATE_KEY.IsNull()) {
		state.SOCIAL_AUTH_SAML_SP_PRIVATE_KEY = orig.SOCIAL_AUTH_SAML_SP_PRIVATE_KEY
	}

	return nil
}

func hookSettingsAuthGoogleOauth2(ctx context.Context, source Source, callee Callee, orig *settingsAuthGoogleOauth2TerraformModel, state *settingsAuthGoogleOauth2TerraformModel) (err error) {
	if source == SourceResource && (state == nil || orig == nil) && (callee == CalleeUpdate || callee == CalleeCreate || callee == CalleeRead) {
		return fmt.Errorf("state and orig required for resource")
	}

	if source == SourceResource && callee == CalleeCreate ||
		(state.SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET.Equal(types.StringValue("$encrypted$")) &&
			(source == SourceResource && (callee == CalleeUpdate || callee == CalleeRead)) &&
			!orig.SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET.IsNull()) {
		state.SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET = orig.SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET
	}

	return nil
}

func hookSettingsAuthLdap(ctx context.Context, source Source, callee Callee, orig *settingsAuthLDAPTerraformModel, state *settingsAuthLDAPTerraformModel) (err error) {
	if source == SourceResource && (state == nil || orig == nil) && (callee == CalleeUpdate || callee == CalleeCreate || callee == CalleeRead) {
		return fmt.Errorf("state and orig required for resource")
	}

	if source == SourceResource && callee == CalleeCreate {
		if len(state.AUTH_LDAP_BIND_PASSWORD.ValueString()) > 0 && !orig.AUTH_LDAP_BIND_PASSWORD.IsNull() {
			state.AUTH_LDAP_BIND_PASSWORD = orig.AUTH_LDAP_BIND_PASSWORD
		}

		if len(state.AUTH_LDAP_1_BIND_PASSWORD.ValueString()) > 0 && !orig.AUTH_LDAP_1_BIND_PASSWORD.IsNull() {
			state.AUTH_LDAP_1_BIND_PASSWORD = orig.AUTH_LDAP_1_BIND_PASSWORD
		}

		if len(state.AUTH_LDAP_2_BIND_PASSWORD.ValueString()) > 0 && !orig.AUTH_LDAP_2_BIND_PASSWORD.IsNull() {
			state.AUTH_LDAP_2_BIND_PASSWORD = orig.AUTH_LDAP_2_BIND_PASSWORD
		}

		if len(state.AUTH_LDAP_3_BIND_PASSWORD.ValueString()) > 0 && !orig.AUTH_LDAP_3_BIND_PASSWORD.IsNull() {
			state.AUTH_LDAP_3_BIND_PASSWORD = orig.AUTH_LDAP_3_BIND_PASSWORD
		}

		if len(state.AUTH_LDAP_4_BIND_PASSWORD.ValueString()) > 0 && !orig.AUTH_LDAP_4_BIND_PASSWORD.IsNull() {
			state.AUTH_LDAP_4_BIND_PASSWORD = orig.AUTH_LDAP_4_BIND_PASSWORD
		}

		if len(state.AUTH_LDAP_5_BIND_PASSWORD.ValueString()) > 0 && !orig.AUTH_LDAP_5_BIND_PASSWORD.IsNull() {
			state.AUTH_LDAP_5_BIND_PASSWORD = orig.AUTH_LDAP_5_BIND_PASSWORD
		}

	} else if source == SourceResource && (callee == CalleeUpdate || callee == CalleeRead) {

		if len(state.AUTH_LDAP_BIND_PASSWORD.ValueString()) > 0 &&
			state.AUTH_LDAP_BIND_PASSWORD.Equal(types.StringValue("$encrypted$")) &&
			!orig.AUTH_LDAP_BIND_PASSWORD.IsNull() {
			state.AUTH_LDAP_BIND_PASSWORD = orig.AUTH_LDAP_BIND_PASSWORD
		}

		if len(state.AUTH_LDAP_1_BIND_PASSWORD.ValueString()) > 0 &&
			state.AUTH_LDAP_1_BIND_PASSWORD.Equal(types.StringValue("$encrypted$")) &&
			!orig.AUTH_LDAP_1_BIND_PASSWORD.IsNull() {
			state.AUTH_LDAP_1_BIND_PASSWORD = orig.AUTH_LDAP_1_BIND_PASSWORD
		}

		if len(state.AUTH_LDAP_2_BIND_PASSWORD.ValueString()) > 0 &&
			state.AUTH_LDAP_2_BIND_PASSWORD.Equal(types.StringValue("$encrypted$")) &&
			!orig.AUTH_LDAP_2_BIND_PASSWORD.IsNull() {
			state.AUTH_LDAP_2_BIND_PASSWORD = orig.AUTH_LDAP_2_BIND_PASSWORD
		}

		if len(state.AUTH_LDAP_3_BIND_PASSWORD.ValueString()) > 0 &&
			state.AUTH_LDAP_3_BIND_PASSWORD.Equal(types.StringValue("$encrypted$")) &&
			!orig.AUTH_LDAP_3_BIND_PASSWORD.IsNull() {
			state.AUTH_LDAP_3_BIND_PASSWORD = orig.AUTH_LDAP_3_BIND_PASSWORD
		}

		if len(state.AUTH_LDAP_4_BIND_PASSWORD.ValueString()) > 0 &&
			state.AUTH_LDAP_4_BIND_PASSWORD.Equal(types.StringValue("$encrypted$")) &&
			!orig.AUTH_LDAP_4_BIND_PASSWORD.IsNull() {
			state.AUTH_LDAP_4_BIND_PASSWORD = orig.AUTH_LDAP_4_BIND_PASSWORD
		}

		if len(state.AUTH_LDAP_5_BIND_PASSWORD.ValueString()) > 0 &&
			state.AUTH_LDAP_5_BIND_PASSWORD.Equal(types.StringValue("$encrypted$")) &&
			!orig.AUTH_LDAP_5_BIND_PASSWORD.IsNull() {
			state.AUTH_LDAP_5_BIND_PASSWORD = orig.AUTH_LDAP_5_BIND_PASSWORD
		}
	}

	return nil
}
