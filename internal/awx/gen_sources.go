package awx

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/ilijamt/terraform-provider-awx/internal/awx/credentials/aws"
	"github.com/ilijamt/terraform-provider-awx/internal/awx/credentials/net"
)

const (
	ApiVersion string = "24.6.1"
)

// DataSources is a helper function to return all defined data sources
func DataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAdHocCommandDataSource,
		NewApplicationDataSource,
		NewConstructedInventoriesDataSource,
		NewConstructedInventoriesObjectRolesDataSource,
		NewCredentialDataSource,
		NewCredentialInputSourceDataSource,
		NewCredentialObjectRolesDataSource,
		NewCredentialTypeDataSource,
		NewExecutionEnvironmentDataSource,
		NewGroupDataSource,
		NewHostDataSource,
		NewHostObjectRolesDataSource,
		NewInstanceGroupDataSource,
		NewInstanceGroupObjectRolesDataSource,
		NewInventoryDataSource,
		NewInventoryObjectRolesDataSource,
		NewInventorySourceDataSource,
		NewJobTemplateDataSource,
		NewJobTemplateObjectRolesDataSource,
		NewLabelDataSource,
		NewMeDataSource,
		NewNotificationTemplateDataSource,
		NewOrganizationDataSource,
		NewOrganizationObjectRolesDataSource,
		NewProjectDataSource,
		NewProjectObjectRolesDataSource,
		NewScheduleDataSource,
		NewSettingsAuthAzureADOauth2DataSource,
		NewSettingsAuthGithubDataSource,
		NewSettingsAuthGithubEnterpriseDataSource,
		NewSettingsAuthGithubEnterpriseOrgDataSource,
		NewSettingsAuthGithubEnterpriseTeamDataSource,
		NewSettingsAuthGithubOrgDataSource,
		NewSettingsAuthGithubTeamDataSource,
		NewSettingsAuthGoogleOauth2DataSource,
		NewSettingsAuthLDAPDataSource,
		NewSettingsAuthSAMLDataSource,
		NewSettingsJobsDataSource,
		NewSettingsMiscAuthenticationDataSource,
		NewSettingsMiscLoggingDataSource,
		NewSettingsMiscSystemDataSource,
		NewSettingsOpenIDConnectDataSource,
		NewSettingsUIDataSource,
		NewTeamDataSource,
		NewTeamObjectRolesDataSource,
		NewTokensDataSource,
		NewUserDataSource,
		NewWorkflowJobTemplateDataSource,
		NewWorkflowJobTemplateObjectRolesDataSource,
	}
}

// Resources is a helper function to return all defined resources
func Resources() []func() resource.Resource {
	return []func() resource.Resource{
		NewAdHocCommandResource,
		NewApplicationResource,
		NewConstructedInventoriesResource,
		NewCredentialInputSourceResource,
		NewCredentialResource,
		NewCredentialTypeResource,
		NewExecutionEnvironmentResource,
		NewGroupResource,
		NewHostAssociateDisassociateGroupResource,
		NewHostResource,
		NewInstanceGroupResource,
		NewInventoryResource,
		NewInventorySourceResource,
		NewJobTemplateAssociateDisassociateCredentialResource,
		NewJobTemplateAssociateDisassociateInstanceGroupResource,
		NewJobTemplateAssociateDisassociateNotificationTemplateResource,
		NewJobTemplateResource,
		NewJobTemplateSurveyResource,
		NewLabelResource,
		NewNotificationTemplateResource,
		NewOrganizationAssociateDisassociateGalaxyCredentialResource,
		NewOrganizationAssociateDisassociateInstanceGroupResource,
		NewOrganizationResource,
		NewProjectResource,
		NewScheduleResource,
		NewSettingsAuthAzureADOauth2Resource,
		NewSettingsAuthGithubEnterpriseOrgResource,
		NewSettingsAuthGithubEnterpriseResource,
		NewSettingsAuthGithubEnterpriseTeamResource,
		NewSettingsAuthGithubOrgResource,
		NewSettingsAuthGithubResource,
		NewSettingsAuthGithubTeamResource,
		NewSettingsAuthGoogleOauth2Resource,
		NewSettingsAuthLDAPResource,
		NewSettingsAuthSAMLResource,
		NewSettingsJobsResource,
		NewSettingsMiscAuthenticationResource,
		NewSettingsMiscLoggingResource,
		NewSettingsMiscSystemResource,
		NewSettingsOpenIDConnectResource,
		NewSettingsUIResource,
		NewTeamAssociateDisassociateRoleResource,
		NewTeamResource,
		NewTokensResource,
		NewUserAssociateDisassociateRoleResource,
		NewUserResource,
		NewWorkflowJobTemplateAssociateDisassociateNotificationTemplateResource,
		NewWorkflowJobTemplateResource,
		NewWorkflowJobTemplateSurveyResource,
		aws.NewResource,
		net.NewResource,
	}
}
