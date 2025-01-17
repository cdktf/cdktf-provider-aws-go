// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gameliftfleet


type GameliftFleetRuntimeConfigurationServerProcess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/gamelift_fleet#concurrent_executions GameliftFleet#concurrent_executions}.
	ConcurrentExecutions *float64 `field:"required" json:"concurrentExecutions" yaml:"concurrentExecutions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/gamelift_fleet#launch_path GameliftFleet#launch_path}.
	LaunchPath *string `field:"required" json:"launchPath" yaml:"launchPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/gamelift_fleet#parameters GameliftFleet#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

