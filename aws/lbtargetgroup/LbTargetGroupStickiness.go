// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbtargetgroup


type LbTargetGroupStickiness struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/lb_target_group#type LbTargetGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/lb_target_group#cookie_duration LbTargetGroup#cookie_duration}.
	CookieDuration *float64 `field:"optional" json:"cookieDuration" yaml:"cookieDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/lb_target_group#cookie_name LbTargetGroup#cookie_name}.
	CookieName *string `field:"optional" json:"cookieName" yaml:"cookieName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/lb_target_group#enabled LbTargetGroup#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

