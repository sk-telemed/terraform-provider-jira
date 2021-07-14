terraform {
  required_providers {
    jira = {
      source  = "sk-telemed/jira"
      version = "1.1.0"
    }
  }
}

provider "jira" {
  url      = "https://test.atlassian.net/"
  user     = "sk-telemed"
  password = "sk-telemed"
}

// This example works for JIRA projects that support epics
data "jira_field" "epic_name" {
  name = "Epic Name"
}

data "jira_field" "epic_link" {
  name = "Epic Link"
}

resource "jira_issue" "example_epic" {
  assignee = "sk-telemed"
  reporter = "sk-telemed"

  issue_type = "Epic"

  // description is optional
  description = "This is an epic description"
  summary     = "This is an epic summary"

  labels = ["one", "two", "buckle-my-shoe"]

  // System and custom fields are optional; see the field data source to reference internal JIRA field IDs by name
  fields = {
    (data.jira_field.epic_name.id) = "Example epic name"
  }

  project_key = data.jira_project.test_project.key
}

resource "jira_issue" "example" {
  assignee = "sk-telemed"
  reporter = "sk-telemed"

  issue_type = "Task"

  // description is optional
  description = "This is a test issue that's part of an epic"
  summary     = "Created using Terraform"
  labels      = ["label1", "label2"]
  fields      = {
    (data.jira_field.epic_link.id) = jira_issue.example_epic.issue_key
  }

  project_key = data.jira_project.test_project.key
}

data "jira_project" "test_project" {
  key = "EPIC"
}

