// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbtargetgroup


type LbTargetGroupTargetFailover struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/lb_target_group#on_deregistration LbTargetGroup#on_deregistration}.
	OnDeregistration *string `field:"required" json:"onDeregistration" yaml:"onDeregistration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/lb_target_group#on_unhealthy LbTargetGroup#on_unhealthy}.
	OnUnhealthy *string `field:"required" json:"onUnhealthy" yaml:"onUnhealthy"`
}

