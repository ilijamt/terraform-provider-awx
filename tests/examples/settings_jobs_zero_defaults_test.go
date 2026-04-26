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

// TestIntegration_SettingsJobsZeroDefaults reproduces a Create-time
// reconciliation bug on awx_settings_jobs. Setting
// event_stdout_max_bytes_display, max_websocket_event_rate, or
// stdout_max_bytes_display to 0 (a valid value per AWX's min_value: 0)
// raises "Provider produced inconsistent result after apply" because
// the generated body-request struct tags these int64 fields with
// json:"...,omitempty", which strips 0 from the PATCH wire. AWX leaves
// the field at its runtime default (1024 / 30 / 1048576) and echoes
// that default back, leaving state inconsistent with the planned 0.
//
// Three steps:
//  1. Apply main.tf with the three problem fields = 0 — exercises the
//     bug. After the fix the values must round-trip as 0 in state.
//  2. Apply update.tf with non-zero values — confirms mutations also
//     flow through cleanly (i.e. the fix didn't break the normal path).
//  3. Re-apply main.tf — confirms round-tripping back to 0 works after
//     the field was previously non-zero.
func TestIntegration_SettingsJobsZeroDefaults(t *testing.T) {
	httpClient := NewVCRClient(t, "settings_jobs_zero_defaults")
	cfg := ReadFixture(t, filepath.Join("settings_jobs_zero_defaults", "main.tf"))
	updated := ReadFixture(t, filepath.Join("settings_jobs_zero_defaults", "update.tf"))

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
					resource.TestCheckResourceAttr("awx_settings_jobs.zero_defaults", "event_stdout_max_bytes_display", "0"),
					resource.TestCheckResourceAttr("awx_settings_jobs.zero_defaults", "max_websocket_event_rate", "0"),
					resource.TestCheckResourceAttr("awx_settings_jobs.zero_defaults", "stdout_max_bytes_display", "0"),
					resource.TestCheckResourceAttr("awx_settings_jobs.zero_defaults", "schedule_max_jobs", "10"),
					resource.TestCheckResourceAttr("awx_settings_jobs.zero_defaults", "max_forks", "200"),
				),
			},
			{
				Config: providerHeader(t) + updated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("awx_settings_jobs.zero_defaults", "event_stdout_max_bytes_display", "2048"),
					resource.TestCheckResourceAttr("awx_settings_jobs.zero_defaults", "max_websocket_event_rate", "60"),
					resource.TestCheckResourceAttr("awx_settings_jobs.zero_defaults", "stdout_max_bytes_display", "524288"),
					resource.TestCheckResourceAttr("awx_settings_jobs.zero_defaults", "schedule_max_jobs", "25"),
					resource.TestCheckResourceAttr("awx_settings_jobs.zero_defaults", "max_forks", "100"),
				),
			},
			{
				Config: providerHeader(t) + cfg,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("awx_settings_jobs.zero_defaults", "event_stdout_max_bytes_display", "0"),
					resource.TestCheckResourceAttr("awx_settings_jobs.zero_defaults", "max_websocket_event_rate", "0"),
					resource.TestCheckResourceAttr("awx_settings_jobs.zero_defaults", "stdout_max_bytes_display", "0"),
					resource.TestCheckResourceAttr("awx_settings_jobs.zero_defaults", "schedule_max_jobs", "10"),
					resource.TestCheckResourceAttr("awx_settings_jobs.zero_defaults", "max_forks", "200"),
				),
			},
		},
	})
}
