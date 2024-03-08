// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbtargetgroup


type LbTargetGroupTargetHealthState struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/lb_target_group#enable_unhealthy_connection_termination LbTargetGroup#enable_unhealthy_connection_termination}.
	EnableUnhealthyConnectionTermination interface{} `field:"required" json:"enableUnhealthyConnectionTermination" yaml:"enableUnhealthyConnectionTermination"`
}

