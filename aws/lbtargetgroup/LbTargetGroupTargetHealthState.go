// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbtargetgroup


type LbTargetGroupTargetHealthState struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/lb_target_group#enable_unhealthy_connection_termination LbTargetGroup#enable_unhealthy_connection_termination}.
	EnableUnhealthyConnectionTermination interface{} `field:"required" json:"enableUnhealthyConnectionTermination" yaml:"enableUnhealthyConnectionTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/lb_target_group#unhealthy_draining_interval LbTargetGroup#unhealthy_draining_interval}.
	UnhealthyDrainingInterval *float64 `field:"optional" json:"unhealthyDrainingInterval" yaml:"unhealthyDrainingInterval"`
}

