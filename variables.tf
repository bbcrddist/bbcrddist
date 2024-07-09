variable "aws_access_key" {
  default = ""
}

variable "aws_secret_key" {
  default = ""
}

variable "region" {
  description = "The AWS region to create resources in."
  default = "ap-southeast-1"
}

variable "environment_name" {
    default = "ECS Instance"
    description = "Environment name to tag EC2 resources with (tag=environment)"
}

variable "ecs_cluster_name" {
  description = "The name of the Amazon ECS cluster."
  default = "default"
}

variable "instance_type" {
  default = "t2.micro"
}

variable "key_name" {
  description = "The aws ssh key name."
  default = "p_singapore"
}

variable "node" {
  description = "Properties for MachinePool node types"
  type = object({
    ctl_plane = map(any)
    worker    = map(any)
  })
}

variable "rancher_env" {
  description = "Variables for Rancher environment"
  type = object({
    cloud_credential    = string
    cluster_annotations = map(string)
    cluster_labels      = map(string)
    rke2_version        = string
  })
}

variable "vsphere_env" {
  description = "Variables for vSphere environment"
  type = object({
    cloud_image_name = string
    datacenter       = string
    datastore        = string
    library_name     = string
    server           = string
    vm_network       = string
  })
}
