// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceDeploymentConfigurationLinearConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_service#step_bake_time_in_minutes EcsService#step_bake_time_in_minutes}.
	StepBakeTimeInMinutes *string `field:"optional" json:"stepBakeTimeInMinutes" yaml:"stepBakeTimeInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_service#step_percent EcsService#step_percent}.
	StepPercent *float64 `field:"optional" json:"stepPercent" yaml:"stepPercent"`
}

