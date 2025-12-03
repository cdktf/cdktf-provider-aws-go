// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alblistenerrule


type AlbListenerRuleTransformUrlRewriteConfigRewrite struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/alb_listener_rule#regex AlbListenerRule#regex}.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/alb_listener_rule#replace AlbListenerRule#replace}.
	Replace *string `field:"required" json:"replace" yaml:"replace"`
}

