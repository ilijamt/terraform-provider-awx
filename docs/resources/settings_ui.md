---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awx_settings_ui Resource - awx"
subcategory: ""
description: |-
  
---

# awx_settings_ui (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `custom_login_info` (String) If needed, you can add specific information (such as a legal notice or a disclaimer) to a text box in the login modal using this setting. Any content added must be in plain text or an HTML fragment, as other markup languages are not supported.
- `custom_logo` (String) To set up a custom logo, provide a file that you create. For the custom logo to look its best, use a .png file with a transparent background. GIF, PNG and JPEG formats are supported.
- `max_ui_job_events` (Number) Maximum number of job events for the UI to retrieve within a single request.
- `ui_live_updates_enabled` (Boolean) If disabled, the page will not refresh when events are received. Reloading the page will be required to get the latest details.

### Read-Only

- `pendo_tracking_state` (String) Enable or Disable User Analytics Tracking.
