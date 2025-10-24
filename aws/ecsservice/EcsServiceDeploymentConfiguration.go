// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceDeploymentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/ecs_service#bake_time_in_minutes EcsService#bake_time_in_minutes}.
	BakeTimeInMinutes *string `field:"optional" json:"bakeTimeInMinutes" yaml:"bakeTimeInMinutes"`
	// lifecycle_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/ecs_service#lifecycle_hook EcsService#lifecycle_hook}
	LifecycleHook interface{} `field:"optional" json:"lifecycleHook" yaml:"lifecycleHook"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/ecs_service#strategy EcsService#strategy}.
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

