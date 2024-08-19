resource "minikube_cluster" "this" {
  vm                 = true
  cluster_name       = var.cluster_name
  nodes              = 1
  kubernetes_version = var.kubernetes_version
  memory             = "8192mb"
  cpus               = 4
  container_runtime  = "containerd"
  addons = [
    "default-storageclass",
    "storage-provisioner",
  ]
}
