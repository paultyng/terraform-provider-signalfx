---
layout: "signalfx"
page_title: "SignalFx: signalfx_opsgenie_integration"
sidebar_current: "docs-signalfx-resource-opsgenie-integration"
description: |-
  Allows Terraform to create and manage SignalFx Opsgenie Integrations
---

# Resource: signalfx_opsgenie_integration

SignalFx Opsgenie integration.

~> **NOTE** When managing integrations you'll need to use a session token for an administrator to authenticate the SignalFx provider. See [Operations that require a session token for an administrator].(https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.

## Example Usage

```tf
resource "signalfx_opsgenie_integration" "opgenie_myteam" {
  name    = "Opsgenie - My Team"
  enabled = true
  api_key = "farts"
  api_url = "https://api.opsgenie.com"
}
```

## Argument Reference

* `name` - (Required) Name of the integration.
* `enabled` - (Required) Whether the integration is enabled.
* `api_key` - (Required) The API key
* `api_url` - (Optional) Opsgenie API URL. Will default to `https://api.opsgenie.com`. You might also want `https://api.eu.opsgenie.com`.

## Attributes Reference

In a addition to all arguments above, the following attributes are exported:

* `id` - The ID of the integration.
