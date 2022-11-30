package ecscluster


type EcsClusterServiceConnectDefaults struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_cluster#namespace EcsCluster#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

