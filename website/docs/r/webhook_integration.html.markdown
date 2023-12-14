---
layout: "signalfx"
page_title: "Splunk Observability Cloud: signalfx_webhook_integration"
sidebar_current: "docs-signalfx-resource-webhook-integration"
description: |-
  Allows Terraform to create and manage webhook integrations for Splunk Observability Cloud
---

# Resource: signalfx_webhook_integration

Splunk Observability Cloud webhook integration.

~> **NOTE** When managing integrations, use a session token of an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.

## Example

```tf
resource "signalfx_webhook_integration" "webhook_myteam" {
  name          = "Webhook - My Team"
  enabled       = true
  url           = "https://www.example.com"
  shared_secret = "abc1234"

  headers {
    header_key   = "some_header"
    header_value = "value_for_that_header"
  }
}
```

## Arguments

* `name` - (Required) Name of the integration.
* `enabled` - (Required) Whether the integration is enabled.
* `url` - (Required) The URL to request
* `shared_secret` - (Optional)
* `headers` - (Optional) A header to send with the request
  * `header_key` - (Required) The key of the header to send
  * `header_value` - (Required) The value of the header to send

## Attributes

In a addition to all arguments above, the following attributes are exported:

* `id` - The ID of the integration.
