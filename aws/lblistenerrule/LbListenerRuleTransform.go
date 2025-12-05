// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistenerrule


type LbListenerRuleTransform struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lb_listener_rule#type LbListenerRule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// host_header_rewrite_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lb_listener_rule#host_header_rewrite_config LbListenerRule#host_header_rewrite_config}
	HostHeaderRewriteConfig *LbListenerRuleTransformHostHeaderRewriteConfig `field:"optional" json:"hostHeaderRewriteConfig" yaml:"hostHeaderRewriteConfig"`
	// url_rewrite_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lb_listener_rule#url_rewrite_config LbListenerRule#url_rewrite_config}
	UrlRewriteConfig *LbListenerRuleTransformUrlRewriteConfig `field:"optional" json:"urlRewriteConfig" yaml:"urlRewriteConfig"`
}

