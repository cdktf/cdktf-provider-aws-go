// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceDeploymentConfigurationLifecycleHook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ecs_service#hook_target_arn EcsService#hook_target_arn}.
	HookTargetArn *string `field:"required" json:"hookTargetArn" yaml:"hookTargetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ecs_service#lifecycle_stages EcsService#lifecycle_stages}.
	LifecycleStages *[]*string `field:"required" json:"lifecycleStages" yaml:"lifecycleStages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ecs_service#role_arn EcsService#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ecs_service#hook_details EcsService#hook_details}.
	HookDetails *string `field:"optional" json:"hookDetails" yaml:"hookDetails"`
}

