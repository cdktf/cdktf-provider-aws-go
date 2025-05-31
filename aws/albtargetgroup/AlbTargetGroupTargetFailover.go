// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package albtargetgroup


type AlbTargetGroupTargetFailover struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/alb_target_group#on_deregistration AlbTargetGroup#on_deregistration}.
	OnDeregistration *string `field:"required" json:"onDeregistration" yaml:"onDeregistration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/alb_target_group#on_unhealthy AlbTargetGroup#on_unhealthy}.
	OnUnhealthy *string `field:"required" json:"onUnhealthy" yaml:"onUnhealthy"`
}

