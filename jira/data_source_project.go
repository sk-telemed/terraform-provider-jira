package jira

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
)

func dataSourceJiraProject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceJiraProjectRead,
		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceJiraProjectRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, _, err := config.jiraClient.Project.Get(d.Get("key").(string))
	if err != nil {
		return errors.Wrap(err, "getting jira project failed")
	}
	d.SetId(project.ID)
	d.Set("project_id", project.ID)
	d.Set("key", project.Key)
	d.Set("name", project.Name)
	d.Set("description", project.Description)
	d.Set("lead", project.Lead)
	d.Set("lead_account_id", project.Lead.AccountID)
	d.Set("url", project.URL)
	d.Set("assignee_type", project.AssigneeType)
	d.Set("category_id", project.ProjectCategory.ID)

	issuesecuritylevelscheme, err := GetJiraResourceID(config.jiraClient, fmt.Sprintf("%s/%s/issuesecuritylevelscheme", projectAPIEndpoint, project.ID))
	if err != nil {
		return errors.Wrap(err, "getting issuesecuritylevelscheme failed")
	}
	d.Set("issue_security_scheme", issuesecuritylevelscheme)

	notificationscheme, err := GetJiraResourceID(config.jiraClient, fmt.Sprintf("%s/%s/notificationscheme", projectAPIEndpoint, project.ID))
	if err != nil {
		return errors.Wrap(err, "getting notificationscheme failed")
	}
	d.Set("notification_scheme", notificationscheme)

	permissionscheme, err := GetJiraResourceID(config.jiraClient, fmt.Sprintf("%s/%s/permissionscheme", projectAPIEndpoint, project.ID))
	if err != nil {
		return errors.Wrap(err, "getting permissionscheme failed")
	}
	d.Set("permission_scheme", permissionscheme)

	if d.Id() == "" {
		return fmt.Errorf("project %q not found", d.Get("name").(string))
	}
	return nil
}
