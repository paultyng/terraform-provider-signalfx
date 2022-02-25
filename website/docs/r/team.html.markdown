---
layout: "signalfx"
page_title: "SignalFx: signalfx_resource"
sidebar_current: "docs-signalfx-resource-team"
description: |-
  Allows Terraform to create and manage SignalFx teams
---

# Resource: signalfx_team

Handles management of SignalFx teams.

You can configure [team notification policies](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html) using this resource and the various `notifications_*` properties.

~> **NOTE** When managing teams use a session token for an administrator to authenticate the SignalFx provider. See [Operations that require a session token for an administrator].(https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator).

## Example Usage

```tf
resource "signalfx_team" "myteam0" {
  name        = "Best Team Ever"
  description = "Super great team no jerks definitely"

  members = [
    "userid1",
    "userid2",
    # …
  ]

  notifications_critical = [
    "PagerDuty,credentialId"
  ]

  notifications_info = [
    "Email,notify@example.com"
  ]
}
```

## Argument Reference

The following arguments are supported in the resource block:

* `name` - (Required) Name of the team.
* `description` - (Optional) Description of the team.
* `members` - (Optional) List of user IDs to include in the team.
* `notifications_critical` - (Optional) Where to send notifications for critical alerts
* `notifications_default` - (Optional) Where to send notifications for default alerts
* `notifications_info` - (Optional) Where to send notifications for info alerts
* `notifications_major` - (Optional) Where to send notifications for major alerts
* `notifications_minor` - (Optional) Where to send notifications for minor alerts
* `notifications_warning` - (Optional) Where to send notifications for warning alerts

## Attributes Reference

In a addition to all arguments above, the following attributes are exported:

* `id` - The ID of the team.
* `url` - The URL of the team.
