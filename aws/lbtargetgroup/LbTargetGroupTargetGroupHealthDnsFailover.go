// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbtargetgroup


type LbTargetGroupTargetGroupHealthDnsFailover struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/lb_target_group#minimum_healthy_targets_count LbTargetGroup#minimum_healthy_targets_count}.
	MinimumHealthyTargetsCount *string `field:"optional" json:"minimumHealthyTargetsCount" yaml:"minimumHealthyTargetsCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/lb_target_group#minimum_healthy_targets_percentage LbTargetGroup#minimum_healthy_targets_percentage}.
	MinimumHealthyTargetsPercentage *string `field:"optional" json:"minimumHealthyTargetsPercentage" yaml:"minimumHealthyTargetsPercentage"`
}

