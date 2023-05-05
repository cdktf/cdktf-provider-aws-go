package ecstaskdefinition


type EcsTaskDefinitionInferenceAccelerator struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/ecs_task_definition#device_name EcsTaskDefinition#device_name}.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/ecs_task_definition#device_type EcsTaskDefinition#device_type}.
	DeviceType *string `field:"required" json:"deviceType" yaml:"deviceType"`
}

