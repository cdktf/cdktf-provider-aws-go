// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alblistenerrule


type AlbListenerRuleTransform struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/alb_listener_rule#type AlbListenerRule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// host_header_rewrite_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/alb_listener_rule#host_header_rewrite_config AlbListenerRule#host_header_rewrite_config}
	HostHeaderRewriteConfig *AlbListenerRuleTransformHostHeaderRewriteConfig `field:"optional" json:"hostHeaderRewriteConfig" yaml:"hostHeaderRewriteConfig"`
	// url_rewrite_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/alb_listener_rule#url_rewrite_config AlbListenerRule#url_rewrite_config}
	UrlRewriteConfig *AlbListenerRuleTransformUrlRewriteConfig `field:"optional" json:"urlRewriteConfig" yaml:"urlRewriteConfig"`
}

