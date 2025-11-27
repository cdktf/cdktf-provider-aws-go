// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceDeploymentConfigurationCanaryConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_service#canary_bake_time_in_minutes EcsService#canary_bake_time_in_minutes}.
	CanaryBakeTimeInMinutes *string `field:"optional" json:"canaryBakeTimeInMinutes" yaml:"canaryBakeTimeInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_service#canary_percent EcsService#canary_percent}.
	CanaryPercent *float64 `field:"optional" json:"canaryPercent" yaml:"canaryPercent"`
}

