package jira

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceJiraProject_basic(t *testing.T) {

	projectName := "HP"
	dataSourceName := "data.jira_project.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceJiraProjectConfig(projectName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "key"),
				),
			},
		},
	})
}

func testAccDataSourceJiraProjectConfig(projectName string) string {
	return fmt.Sprintf(`
data "jira_project" "test" {
  key = "%s"
}`, projectName)
}
