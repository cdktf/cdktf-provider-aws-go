package ecstaskdefinition


type EcsTaskDefinitionVolumeDockerVolumeConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/ecs_task_definition#autoprovision EcsTaskDefinition#autoprovision}.
	Autoprovision interface{} `field:"optional" json:"autoprovision" yaml:"autoprovision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/ecs_task_definition#driver EcsTaskDefinition#driver}.
	Driver *string `field:"optional" json:"driver" yaml:"driver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/ecs_task_definition#driver_opts EcsTaskDefinition#driver_opts}.
	DriverOpts *map[string]*string `field:"optional" json:"driverOpts" yaml:"driverOpts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/ecs_task_definition#labels EcsTaskDefinition#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/ecs_task_definition#scope EcsTaskDefinition#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

