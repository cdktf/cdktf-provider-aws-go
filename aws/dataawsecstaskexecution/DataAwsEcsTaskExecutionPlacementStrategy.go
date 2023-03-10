package dataawsecstaskexecution


type DataAwsEcsTaskExecutionPlacementStrategy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ecs_task_execution#type DataAwsEcsTaskExecution#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ecs_task_execution#field DataAwsEcsTaskExecution#field}.
	Field *string `field:"optional" json:"field" yaml:"field"`
}

