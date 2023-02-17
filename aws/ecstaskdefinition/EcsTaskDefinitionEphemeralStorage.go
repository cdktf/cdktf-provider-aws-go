package ecstaskdefinition


type EcsTaskDefinitionEphemeralStorage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_task_definition#size_in_gib EcsTaskDefinition#size_in_gib}.
	SizeInGib *float64 `field:"required" json:"sizeInGib" yaml:"sizeInGib"`
}

