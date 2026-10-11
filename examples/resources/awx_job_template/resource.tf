resource "awx_job_template" "instance_group" {
  name               = "Job Demo Instance Group"
  inventory          = awx_inventory.demo_inventory.id
  verbosity          = 0
  job_type           = "run"
  playbook           = "hello_world.yml"
  project            = awx_project.demo_project.id
  scm_branch         = ""
  job_slice_count    = 2
  allow_simultaneous = true
}
