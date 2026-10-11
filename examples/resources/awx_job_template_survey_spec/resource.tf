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
