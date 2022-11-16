terraform {
  required_providers {
    awx = {
      source = "registry.terraform.io/ilijamt/awx"
    }
  }
}

provider "awx" {}

resource "awx_organization" "demo_organization" {
  name                = "Job Template Organization"
}

resource "awx_project" "demo_project" {
  name         = "Job Demo Project"
  organization = awx_organization.demo_organization.id
  scm_url      = "https://github.com/ansible/ansible-tower-samples"
  scm_type     = "git"
  scm_clean    = false
}

resource "awx_inventory" "demo_inventory" {
  name         = "Job Demo Inventory"
  organization = awx_organization.demo_organization.id
}

resource "awx_job_template" "demo_job_template" {
  name               = "Job Demo Job Template"
  inventory          = awx_inventory.demo_inventory.id
  verbosity          = 0
  job_type           = "run"
  playbook           = "hello_world.yml"
  project            = awx_project.demo_project.id
  scm_branch         = ""
  job_slice_count    = 2
  allow_simultaneous = true
}

 resource "awx_job_template_survey_spec" "demo_job_template" {
   job_template_id = awx_job_template.demo_job_template.id
   spec = jsonencode([{
     type                 = "float"
     max                  = 1024
     min                  = 5
     choices              = ""
     default              = 10
     question_description = "I do not know"
     new_question         = false
     question_name        = "What is the percentage of failure?"
     required             = true
     variable             = "pct_failure"
   }])
 }
