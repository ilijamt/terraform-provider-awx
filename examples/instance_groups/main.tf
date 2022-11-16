terraform {
  required_providers {
    awx = {
      source = "registry.terraform.io/ilijamt/awx"
    }
  }
}

provider "awx" {}

resource "awx_instance_group" "ig" {
  name               = "Demo Instance Group"
  is_container_group = true
  pod_spec_override = jsonencode({
    "apiVersion" : "v1",
    "kind" : "Pod",
    "metadata" : {
      "namespace" : "awx"
    },
    "spec" : {
      "serviceAccountName" : "default",
      "automountServiceAccountToken" : false,
      "containers" : [
        {
          "image" : "quay.io/ansible/awx-ee:latest",
          "name" : "worker",
          "args" : [
            "ansible-runner",
            "worker",
            "--private-data-dir=/runner"
          ],
          "resources" : {
            "requests" : {
              "cpu" : "250m",
              "memory" : "100Mi"
            }
          }
        }
      ]
    }
  })
}