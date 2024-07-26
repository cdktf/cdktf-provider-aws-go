// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeTargetParametersEcsTaskParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/pipes_pipe#task_definition_arn PipesPipe#task_definition_arn}.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// capacity_provider_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/pipes_pipe#capacity_provider_strategy PipesPipe#capacity_provider_strategy}
	CapacityProviderStrategy interface{} `field:"optional" json:"capacityProviderStrategy" yaml:"capacityProviderStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/pipes_pipe#enable_ecs_managed_tags PipesPipe#enable_ecs_managed_tags}.
	EnableEcsManagedTags interface{} `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/pipes_pipe#enable_execute_command PipesPipe#enable_execute_command}.
	EnableExecuteCommand interface{} `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/pipes_pipe#group PipesPipe#group}.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/pipes_pipe#launch_type PipesPipe#launch_type}.
	LaunchType *string `field:"optional" json:"launchType" yaml:"launchType"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/pipes_pipe#network_configuration PipesPipe#network_configuration}
	NetworkConfiguration *PipesPipeTargetParametersEcsTaskParametersNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/pipes_pipe#overrides PipesPipe#overrides}
	Overrides *PipesPipeTargetParametersEcsTaskParametersOverrides `field:"optional" json:"overrides" yaml:"overrides"`
	// placement_constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/pipes_pipe#placement_constraint PipesPipe#placement_constraint}
	PlacementConstraint interface{} `field:"optional" json:"placementConstraint" yaml:"placementConstraint"`
	// placement_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/pipes_pipe#placement_strategy PipesPipe#placement_strategy}
	PlacementStrategy interface{} `field:"optional" json:"placementStrategy" yaml:"placementStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/pipes_pipe#platform_version PipesPipe#platform_version}.
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/pipes_pipe#propagate_tags PipesPipe#propagate_tags}.
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/pipes_pipe#reference_id PipesPipe#reference_id}.
	ReferenceId *string `field:"optional" json:"referenceId" yaml:"referenceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/pipes_pipe#tags PipesPipe#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/pipes_pipe#task_count PipesPipe#task_count}.
	TaskCount *float64 `field:"optional" json:"taskCount" yaml:"taskCount"`
}

