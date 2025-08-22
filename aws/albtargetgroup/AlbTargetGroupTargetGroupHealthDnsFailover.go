// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package albtargetgroup


type AlbTargetGroupTargetGroupHealthDnsFailover struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/alb_target_group#minimum_healthy_targets_count AlbTargetGroup#minimum_healthy_targets_count}.
	MinimumHealthyTargetsCount *string `field:"optional" json:"minimumHealthyTargetsCount" yaml:"minimumHealthyTargetsCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/alb_target_group#minimum_healthy_targets_percentage AlbTargetGroup#minimum_healthy_targets_percentage}.
	MinimumHealthyTargetsPercentage *string `field:"optional" json:"minimumHealthyTargetsPercentage" yaml:"minimumHealthyTargetsPercentage"`
}

