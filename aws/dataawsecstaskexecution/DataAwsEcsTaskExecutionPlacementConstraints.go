package dataawsecstaskexecution


type DataAwsEcsTaskExecutionPlacementConstraints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/data-sources/ecs_task_execution#type DataAwsEcsTaskExecution#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/data-sources/ecs_task_execution#expression DataAwsEcsTaskExecution#expression}.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

