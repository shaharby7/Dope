locals {
  environment_vars = read_terragrunt_config(find_in_parent_folders("env.hcl"))
  global_vars      = read_terragrunt_config(find_in_parent_folders("globals.hcl"))
  stage            = local.environment_vars.locals.environment
  charts_path      = local.global_vars.locals.charts_path
}

terraform {
  source = "../../../modules//minikube"
}

include {
  path = find_in_parent_folders()
}

inputs = {
  cluster_name = "dopeexample"
}