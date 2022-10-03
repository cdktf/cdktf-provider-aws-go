package ecstaskdefinition


type EcsTaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_task_definition#access_point_id EcsTaskDefinition#access_point_id}.
	AccessPointId *string `field:"optional" json:"accessPointId" yaml:"accessPointId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_task_definition#iam EcsTaskDefinition#iam}.
	Iam *string `field:"optional" json:"iam" yaml:"iam"`
}

