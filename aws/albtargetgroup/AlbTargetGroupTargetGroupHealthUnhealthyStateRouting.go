// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package albtargetgroup


type AlbTargetGroupTargetGroupHealthUnhealthyStateRouting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/alb_target_group#minimum_healthy_targets_count AlbTargetGroup#minimum_healthy_targets_count}.
	MinimumHealthyTargetsCount *float64 `field:"optional" json:"minimumHealthyTargetsCount" yaml:"minimumHealthyTargetsCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/alb_target_group#minimum_healthy_targets_percentage AlbTargetGroup#minimum_healthy_targets_percentage}.
	MinimumHealthyTargetsPercentage *string `field:"optional" json:"minimumHealthyTargetsPercentage" yaml:"minimumHealthyTargetsPercentage"`
}

