// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistenerrule


type LbListenerRuleTransformHostHeaderRewriteConfigRewrite struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lb_listener_rule#regex LbListenerRule#regex}.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lb_listener_rule#replace LbListenerRule#replace}.
	Replace *string `field:"required" json:"replace" yaml:"replace"`
}

