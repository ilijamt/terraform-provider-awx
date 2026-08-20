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

// Guards issue #172: the untyped awx_notification_template answered a create
// carrying a secret with "inconsistent values for sensitive attribute", because
// AWX echoes `$encrypted$` inside the JSON blob and injects its own defaults.
func TestIntegration_NotificationTemplates(t *testing.T) {
	httpClient := NewVCRClient(t, "notification_templates")
	cfg := ReadFixture(t, filepath.Join("notification_templates", "main.tf"))
	updated := ReadFixture(t, filepath.Join("notification_templates", "update.tf"))

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
					resource.TestCheckResourceAttr("awx_notification_template_awssns.smoke", "notification_type", "awssns"),
					resource.TestCheckResourceAttr("awx_notification_template_webhook.smoke", "notification_type", "webhook"),
					resource.TestCheckResourceAttrPair("awx_notification_template_slack.smoke", "organization", "awx_organization.org", "id"),

					resource.TestCheckResourceAttr("awx_notification_template_awssns.smoke", "aws_region", "eu-west-1"),
					resource.TestCheckResourceAttr("awx_notification_template_awssns.smoke", "sns_topic_arn", ""),

					resource.TestCheckResourceAttr("awx_notification_template_email.smoke", "port", "587"),
					resource.TestCheckResourceAttr("awx_notification_template_email.smoke", "timeout", "30"),
					resource.TestCheckResourceAttr("awx_notification_template_email.smoke", "use_tls", "true"),
					resource.TestCheckResourceAttr("awx_notification_template_email.smoke", "recipients.#", "2"),
					resource.TestCheckResourceAttr("awx_notification_template_email.smoke", "recipients.0", "ops@example.com"),
					resource.TestCheckResourceAttr("awx_notification_template_email.smoke", "password", "email-secret"),

					resource.TestCheckResourceAttr("awx_notification_template_grafana.smoke", "grafana_key", "grafana-secret"),

					resource.TestCheckResourceAttr("awx_notification_template_irc.smoke", "port", "6697"),
					resource.TestCheckResourceAttr("awx_notification_template_irc.smoke", "targets.0", "#ops"),
					resource.TestCheckResourceAttr("awx_notification_template_irc.smoke", "targets.1", "#alerts"),
					resource.TestCheckResourceAttr("awx_notification_template_irc.smoke", "password", "irc-secret"),

					resource.TestCheckResourceAttr("awx_notification_template_mattermost.smoke", "mattermost_no_verify_ssl", "false"),
					resource.TestCheckResourceAttr("awx_notification_template_rocketchat.smoke", "rocketchat_no_verify_ssl", "false"),

					resource.TestCheckResourceAttr("awx_notification_template_pagerduty.smoke", "client_name", "awx"),
					resource.TestCheckResourceAttr("awx_notification_template_pagerduty.smoke", "token", "pagerduty-secret"),

					resource.TestCheckResourceAttr("awx_notification_template_slack.smoke", "channels.#", "2"),
					resource.TestCheckResourceAttr("awx_notification_template_slack.smoke", "token", "slack-secret"),

					resource.TestCheckResourceAttr("awx_notification_template_twilio.smoke", "to_numbers.#", "2"),
					resource.TestCheckResourceAttr("awx_notification_template_twilio.smoke", "account_token", "twilio-secret"),

					resource.TestCheckResourceAttr("awx_notification_template_webhook.smoke", "url", "https://hooks.example.com/awx"),
					resource.TestCheckResourceAttr("awx_notification_template_webhook.smoke", "headers.%", "0"),
					resource.TestCheckResourceAttr("awx_notification_template_webhook.smoke", "http_method", "POST"),
					resource.TestCheckResourceAttr("awx_notification_template_webhook.smoke", "disable_ssl_verification", "false"),
					resource.TestCheckResourceAttr("awx_notification_template_webhook.smoke", "password", "webhook-secret"),

					resource.TestCheckResourceAttrPair("data.awx_notification_template_webhook.webhook_by_name", "id", "awx_notification_template_webhook.smoke", "id"),
					resource.TestCheckResourceAttr("data.awx_notification_template_irc.irc_by_name", "port", "6697"),
					resource.TestCheckResourceAttr("data.awx_notification_template_irc.irc_by_name", "targets.#", "2"),
				),
			},
			{
				Config: providerHeader(t) + updated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("awx_notification_template_awssns.smoke", "aws_region", "us-east-1"),
					resource.TestCheckResourceAttr("awx_notification_template_awssns.smoke", "aws_secret_access_key", "awssns-secret"),

					resource.TestCheckResourceAttr("awx_notification_template_email.smoke", "port", "465"),
					resource.TestCheckResourceAttr("awx_notification_template_email.smoke", "timeout", "60"),
					resource.TestCheckResourceAttr("awx_notification_template_email.smoke", "use_tls", "false"),
					resource.TestCheckResourceAttr("awx_notification_template_email.smoke", "recipients.#", "1"),
					resource.TestCheckResourceAttr("awx_notification_template_email.smoke", "password", "email-secret-rotated"),

					resource.TestCheckResourceAttr("awx_notification_template_grafana.smoke", "grafana_key", "grafana-secret-rotated"),

					resource.TestCheckResourceAttr("awx_notification_template_irc.smoke", "port", "6667"),
					resource.TestCheckResourceAttr("awx_notification_template_irc.smoke", "use_ssl", "false"),
					resource.TestCheckResourceAttr("awx_notification_template_irc.smoke", "targets.#", "1"),
					resource.TestCheckResourceAttr("awx_notification_template_irc.smoke", "password", "irc-secret-rotated"),

					resource.TestCheckResourceAttr("awx_notification_template_mattermost.smoke", "mattermost_no_verify_ssl", "true"),
					resource.TestCheckResourceAttr("awx_notification_template_rocketchat.smoke", "rocketchat_no_verify_ssl", "true"),

					resource.TestCheckResourceAttr("awx_notification_template_pagerduty.smoke", "client_name", "awx-prod"),
					resource.TestCheckResourceAttr("awx_notification_template_pagerduty.smoke", "token", "pagerduty-secret-rotated"),

					resource.TestCheckResourceAttr("awx_notification_template_slack.smoke", "channels.#", "3"),
					resource.TestCheckResourceAttr("awx_notification_template_slack.smoke", "channels.2", "#incidents"),
					resource.TestCheckResourceAttr("awx_notification_template_slack.smoke", "token", "slack-secret-rotated"),

					resource.TestCheckResourceAttr("awx_notification_template_twilio.smoke", "to_numbers.#", "1"),
					resource.TestCheckResourceAttr("awx_notification_template_twilio.smoke", "account_token", "twilio-secret-rotated"),

					resource.TestCheckResourceAttr("awx_notification_template_webhook.smoke", "url", "https://hooks.example.com/awx-v2"),
					resource.TestCheckResourceAttr("awx_notification_template_webhook.smoke", "headers.X-Source", "awx"),
					resource.TestCheckResourceAttr("awx_notification_template_webhook.smoke", "http_method", "PUT"),
					resource.TestCheckResourceAttr("awx_notification_template_webhook.smoke", "disable_ssl_verification", "true"),
					resource.TestCheckResourceAttr("awx_notification_template_webhook.smoke", "password", "webhook-secret-rotated"),
				),
			},
			{
				// AWX returns `$encrypted$`; the merge hook only restores a
				// secret when prior plan state exists, which import lacks.
				ResourceName:            "awx_notification_template_webhook.smoke",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
			{
				// rocketchat has no secret fields, so this excludes nothing.
				ResourceName:      "awx_notification_template_rocketchat.smoke",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:            "awx_notification_template_irc.smoke",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
		},
	})
}
