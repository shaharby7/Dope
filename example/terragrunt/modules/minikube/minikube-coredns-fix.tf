# ## Ugly fix of the minikube coredns file
# 
# 
# # Wait for coredns do be created normally
# resource "time_sleep" "wait_problematic_coredns" {
#   depends_on = [minikube_cluster.this]
# 
#   create_duration = "60s"
# }
# 
# # First - delete the problematic config map
# resource "null_resource" "delete_problematic_cm" {
#   depends_on = [ time_sleep.wait_problematic_coredns ]
# 
#   provisioner "local-exec" {
#     command = <<-EOC
# kubectl config set-context ${var.cluster_name}
# kubectl delete cm coredns -n kube-system 
# EOC
#   }
# }
# 
# # Recreate the configmap correctly 
# resource "kubernetes_config_map" "coredns" {
#   depends_on = [ null_resource.delete_problematic_cm ]
#   metadata {
#     name = "coredns"
#     namespace = "kube-system"
#   }
#   data = {
#     Corefile = "${file("${path.module}/data/coredns-config")}"
#   }
# }
# 
# # Restart the pod of the coredns to re-load the correct config
# resource "null_resource" "restart-coredns-pod" {
#   depends_on = [ kubernetes_config_map.coredns ]
# 
#   provisioner "local-exec" {
#     command = <<-EOC
# kubectl config set-context ${var.cluster_name}
# kubectl rollout restart deployment/coredns -n kube-system
# EOC
#   }
# }