//go:build integration

package examples

import (
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/awx"
	"github.com/ilijamt/terraform-provider-awx/internal/provider"
	"github.com/ilijamt/terraform-provider-awx/version"
)

// TestIntegration_SettingsMisc covers the singleton settings resources from
// examples/settings — except awx_settings_jobs, which is exercised by
// TestIntegration_SettingsJobsZeroDefaults.
//
// Each settings_* resource is a singleton (PATCHes a single AWX record). The
// test proves that the typical mix of scalar, list, and JSON-encoded fields
// across awx_settings_ui, awx_settings_oidc, awx_settings_misc_logging,
// awx_settings_misc_system, and awx_settings_misc_authentication round-trips
// through both create and update without drift, and that the corresponding
// awx_settings_misc_system data source reads the same values back.
//
// No import step: settings resources are singletons identified by a fixed
// path, and import semantics are uninteresting compared to the round-trip
// behavior verified above.
func TestIntegration_SettingsMisc(t *testing.T) {
	httpClient := NewVCRClient(t, "settings_misc")
	cfg := ReadFixture(t, filepath.Join("settings_misc", "main.tf"))
	updated := ReadFixture(t, filepath.Join("settings_misc", "update.tf"))

	factories := map[string]func() (tfprotov6.ProviderServer, error){
		"awx": providerserver.NewProtocol6WithError(
			provider.NewFuncProvider(version.Version, httpClient, awx.Resources(), awx.DataSources())(),
		),
	}

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: factories,
		Steps: []resource.TestStep{
			{
				Config: providerHeader(t) + cfg,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("awx_settings_ui.default", "ui_live_updates_enabled", "true"),
					resource.TestCheckResourceAttr("awx_settings_ui.default", "max_ui_job_events", "2000"),

					resource.TestCheckResourceAttr("awx_settings_oidc.default", "social_auth_oidc_verify_ssl", "false"),

					resource.TestCheckResourceAttr("awx_settings_misc_logging.default", "log_aggregator_level", "INFO"),
					resource.TestCheckResourceAttr("awx_settings_misc_logging.default", "log_aggregator_action_max_disk_usage_gb", "1"),
					resource.TestCheckResourceAttr("awx_settings_misc_logging.default", "log_aggregator_loggers.#", "4"),

					resource.TestCheckResourceAttr("awx_settings_misc_system.default", "tower_url_base", "http://awx.local"),
					resource.TestCheckResourceAttr("awx_settings_misc_system.default", "automation_analytics_gather_interval", "14400"),
					resource.TestCheckResourceAttr("awx_settings_misc_system.default", "automation_analytics_url", "https://example.com"),
					resource.TestCheckResourceAttr("awx_settings_misc_system.default", "remote_host_headers.#", "3"),
					resource.TestCheckResourceAttrPair("awx_settings_misc_system.default", "default_execution_environment", "data.awx_execution_environment.latest", "id"),

					resource.TestCheckResourceAttr("awx_settings_misc_authentication.default", "auth_basic_enabled", "true"),
					resource.TestCheckResourceAttr("awx_settings_misc_authentication.default", "allow_oauth2_for_external_users", "false"),
					resource.TestCheckResourceAttr("awx_settings_misc_authentication.default", "session_cookie_age", "1800000"),

					resource.TestCheckResourceAttrPair("data.awx_settings_misc_system.default", "tower_url_base", "awx_settings_misc_system.default", "tower_url_base"),
				),
			},
			{
				Config: providerHeader(t) + updated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("awx_settings_ui.default", "ui_live_updates_enabled", "false"),
					resource.TestCheckResourceAttr("awx_settings_ui.default", "max_ui_job_events", "4000"),
					resource.TestCheckResourceAttr("awx_settings_ui.default", "custom_login_info", "Welcome to AWX"),

					resource.TestCheckResourceAttr("awx_settings_oidc.default", "social_auth_oidc_verify_ssl", "true"),

					resource.TestCheckResourceAttr("awx_settings_misc_logging.default", "log_aggregator_level", "WARNING"),
					resource.TestCheckResourceAttr("awx_settings_misc_logging.default", "log_aggregator_individual_facts", "true"),
					resource.TestCheckResourceAttr("awx_settings_misc_logging.default", "log_aggregator_loggers.#", "3"),
					resource.TestCheckResourceAttr("awx_settings_misc_logging.default", "log_aggregator_tcp_timeout", "10"),

					resource.TestCheckResourceAttr("awx_settings_misc_system.default", "automation_analytics_gather_interval", "28800"),
					resource.TestCheckResourceAttr("awx_settings_misc_system.default", "automation_analytics_url", "https://updated.example.com"),
					resource.TestCheckResourceAttr("awx_settings_misc_system.default", "activity_stream_enabled_for_inventory_sync", "true"),
					resource.TestCheckResourceAttr("awx_settings_misc_system.default", "remote_host_headers.#", "2"),

					resource.TestCheckResourceAttr("awx_settings_misc_authentication.default", "allow_oauth2_for_external_users", "true"),
					resource.TestCheckResourceAttr("awx_settings_misc_authentication.default", "session_cookie_age", "3600"),
				),
			},
		},
	})
}
