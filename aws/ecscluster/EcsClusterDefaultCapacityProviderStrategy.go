package ecscluster


type EcsClusterDefaultCapacityProviderStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/ecs_cluster#capacity_provider EcsCluster#capacity_provider}.
	CapacityProvider *string `field:"required" json:"capacityProvider" yaml:"capacityProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/ecs_cluster#base EcsCluster#base}.
	Base *float64 `field:"optional" json:"base" yaml:"base"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/ecs_cluster#weight EcsCluster#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

