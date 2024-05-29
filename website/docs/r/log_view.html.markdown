---
layout: "signalfx"
page_title: "Splunk Observability Cloud: signalfx_log_view"
sidebar_current: "docs-signalfx-resource-log-view"
description: |-
  Allows Terraform to create and manage log views in Splunk Observability Cloud
---

# Resource: signalfx_log_view

You can add logs data to your Observability Cloud dashboards without turning your logs into metrics first.

A log view displays log lines in a table form in a dashboard and shows you in detail what is happening and why.

## Example

```tf
resource "signalfx_log_view" "my_log_view" {
  name        = "Sample Log View"
  description = "Lorem ipsum dolor sit amet, laudem tibique iracundia at mea. Nam posse dolores ex, nec cu adhuc putent honestatis"

  program_text = <<-EOF
  logs(filter=field('message') == 'Transaction processed' and field('service.name') == 'paymentservice').publish()
  EOF

  time_range = 900
  sort_options {
    descending= false
    field= "severity"
   }

  columns {
        name="severity"
    }
  columns {
        name="time"
    }
  columns {
        name="amount.currency_code"
    }
  columns {
        name="amount.nanos"
    }
  columns {
        name="amount.units"
    }
  columns {
        name="message"
    }

}
```

## Arguments

The following arguments are supported in the resource block:

* `name` - (Required) Name of the log view.
* `program_text` - (Required) Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
* `description` - (Optional) Description of the log view.
* `time_range` - (Optional) From when to display data. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
* `start_time` - (Optional) Seconds since epoch. Used for visualization. Conflicts with `time_range`.
* `end_time` - (Optional) Seconds since epoch. Used for visualization. Conflicts with `time_range`.
* `columns` - (Optional) The column headers to show on the log view.
* `sort_options` - (Optional) The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
* `default_connection` - (Optional) The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.

## Attributes

In a addition to all arguments above, the following attributes are exported:

* `id` - The ID of the log view.
* `url` - The URL of the log view.
