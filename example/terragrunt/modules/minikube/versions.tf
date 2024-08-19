terraform {
  required_providers {
    minikube = {
      source  = "scott-the-programmer/minikube"
      version = ">= 0.3.2"
    }
  }
  required_version = ">= 0.13"
}
