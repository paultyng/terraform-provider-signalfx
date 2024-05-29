package signalfx

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const newLogViewConfig = `
resource "signalfx_log_view" "mychart0" {
  name = "Chart Name"
  description = "Chart Description"
  program_text = "logs(index=['history','main','o11yhipster','splunklogger','summary']).publish()"

  time_range = 900
  default_connection = "Cosmicbat"
    columns {
      name= "severity"
    }
    columns {
        name= "time"
    }
    columns {
      name= "_raw"
    }
  sort_options {
    descending= false
    field= "severity"
   }
}
`

const updatedLogViewConfig = `
resource "signalfx_log_view" "mychart0" {
  name = "Chart Name NEW"
  description = "Chart Description NEW"
  program_text = "logs().publish()"

  start_time = 1657647022
  end_time = 1657648042

    columns {
      name= "severity"
    }
    columns {
        name= "time"
    }
    columns {
      name= "_raw"
    }
  sort_options {
    descending= true
    field= "severity"
   }
}
`

func TestAccCreateUpdateLogView(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccLogViewDestroy,
		Steps: []resource.TestStep{
			// Create It
			{
				Config: newLogViewConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLogViewResourceExists,
					resource.TestCheckResourceAttr("signalfx_log_view.mychart0", "name", "Chart Name"),
					resource.TestCheckResourceAttr("signalfx_log_view.mychart0", "description", "Chart Description"),
					resource.TestCheckResourceAttr("signalfx_log_view.mychart0", "default_connection", "Cosmicbat"),
					resource.TestCheckResourceAttr("signalfx_log_view.mychart0", "columns.#", "3"),
					resource.TestCheckResourceAttr("signalfx_log_view.mychart0", "columns.0.name", "severity"),
					resource.TestCheckResourceAttr("signalfx_log_view.mychart0", "columns.1.name", "time"),
					resource.TestCheckResourceAttr("signalfx_log_view.mychart0", "columns.2.name", "_raw"),
					resource.TestCheckResourceAttr("signalfx_log_view.mychart0", "sort_options.#", "1"),
					resource.TestCheckResourceAttr("signalfx_log_view.mychart0", "sort_options.0.descending", "false"),
					resource.TestCheckResourceAttr("signalfx_log_view.mychart0", "sort_options.0.field", "severity"),
					resource.TestCheckResourceAttr("signalfx_log_view.mychart0", "time_range", "900"),
				),
			},
			{
				ResourceName:      "signalfx_log_view.mychart0",
				ImportState:       true,
				ImportStateIdFunc: testAccStateIdFunc("signalfx_log_view.mychart0"),
				ImportStateVerify: true,
			},
			// Update Everything
			{
				Config: updatedLogViewConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLogViewResourceExists,
					resource.TestCheckResourceAttr("signalfx_log_view.mychart0", "name", "Chart Name NEW"),
					resource.TestCheckResourceAttr("signalfx_log_view.mychart0", "description", "Chart Description NEW"),
					resource.TestCheckResourceAttr("signalfx_log_view.mychart0", "start_time", "1657647022"),
					resource.TestCheckResourceAttr("signalfx_log_view.mychart0", "end_time", "1657648042"),
					resource.TestCheckResourceAttr("signalfx_log_view.mychart0", "sort_options.0.descending", "true"),
				),
			},
		},
	})
}

func testAccCheckLogViewResourceExists(s *terraform.State) error {
	client := newTestClient()

	for _, rs := range s.RootModule().Resources {
		switch rs.Type {
		case "signalfx_log_view":
			chart, err := client.GetChart(context.TODO(), rs.Primary.ID)
			if chart.Id != rs.Primary.ID || err != nil {
				return fmt.Errorf("Error finding chart %s: %s", rs.Primary.ID, err)
			}
		default:
			return fmt.Errorf("Unexpected resource of type: %s", rs.Type)
		}
	}
	return nil
}

func testAccLogViewDestroy(s *terraform.State) error {
	client := newTestClient()
	for _, rs := range s.RootModule().Resources {
		switch rs.Type {
		case "signalfx_log_view":
			chart, _ := client.GetChart(context.TODO(), rs.Primary.ID)
			if chart != nil {
				return fmt.Errorf("Found deleted chart %s", rs.Primary.ID)
			}
		default:
			return fmt.Errorf("Unexpected resource of type: %s", rs.Type)
		}
	}

	return nil
}
