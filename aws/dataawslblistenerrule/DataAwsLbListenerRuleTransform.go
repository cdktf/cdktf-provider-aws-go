// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawslblistenerrule


type DataAwsLbListenerRuleTransform struct {
	// host_header_rewrite_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/data-sources/lb_listener_rule#host_header_rewrite_config DataAwsLbListenerRule#host_header_rewrite_config}
	HostHeaderRewriteConfig interface{} `field:"optional" json:"hostHeaderRewriteConfig" yaml:"hostHeaderRewriteConfig"`
	// url_rewrite_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/data-sources/lb_listener_rule#url_rewrite_config DataAwsLbListenerRule#url_rewrite_config}
	UrlRewriteConfig interface{} `field:"optional" json:"urlRewriteConfig" yaml:"urlRewriteConfig"`
}

