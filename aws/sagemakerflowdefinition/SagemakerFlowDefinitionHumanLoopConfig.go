// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerflowdefinition


type SagemakerFlowDefinitionHumanLoopConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/sagemaker_flow_definition#human_task_ui_arn SagemakerFlowDefinition#human_task_ui_arn}.
	HumanTaskUiArn *string `field:"required" json:"humanTaskUiArn" yaml:"humanTaskUiArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/sagemaker_flow_definition#task_count SagemakerFlowDefinition#task_count}.
	TaskCount *float64 `field:"required" json:"taskCount" yaml:"taskCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/sagemaker_flow_definition#task_description SagemakerFlowDefinition#task_description}.
	TaskDescription *string `field:"required" json:"taskDescription" yaml:"taskDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/sagemaker_flow_definition#task_title SagemakerFlowDefinition#task_title}.
	TaskTitle *string `field:"required" json:"taskTitle" yaml:"taskTitle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/sagemaker_flow_definition#workteam_arn SagemakerFlowDefinition#workteam_arn}.
	WorkteamArn *string `field:"required" json:"workteamArn" yaml:"workteamArn"`
	// public_workforce_task_price block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/sagemaker_flow_definition#public_workforce_task_price SagemakerFlowDefinition#public_workforce_task_price}
	PublicWorkforceTaskPrice *SagemakerFlowDefinitionHumanLoopConfigPublicWorkforceTaskPrice `field:"optional" json:"publicWorkforceTaskPrice" yaml:"publicWorkforceTaskPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/sagemaker_flow_definition#task_availability_lifetime_in_seconds SagemakerFlowDefinition#task_availability_lifetime_in_seconds}.
	TaskAvailabilityLifetimeInSeconds *float64 `field:"optional" json:"taskAvailabilityLifetimeInSeconds" yaml:"taskAvailabilityLifetimeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/sagemaker_flow_definition#task_keywords SagemakerFlowDefinition#task_keywords}.
	TaskKeywords *[]*string `field:"optional" json:"taskKeywords" yaml:"taskKeywords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/sagemaker_flow_definition#task_time_limit_in_seconds SagemakerFlowDefinition#task_time_limit_in_seconds}.
	TaskTimeLimitInSeconds *float64 `field:"optional" json:"taskTimeLimitInSeconds" yaml:"taskTimeLimitInSeconds"`
}

