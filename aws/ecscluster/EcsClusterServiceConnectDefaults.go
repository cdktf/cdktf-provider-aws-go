package ecscluster


type EcsClusterServiceConnectDefaults struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/ecs_cluster#namespace EcsCluster#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

