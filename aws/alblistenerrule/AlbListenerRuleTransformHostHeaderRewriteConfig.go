// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alblistenerrule


type AlbListenerRuleTransformHostHeaderRewriteConfig struct {
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/alb_listener_rule#rewrite AlbListenerRule#rewrite}
	Rewrite *AlbListenerRuleTransformHostHeaderRewriteConfigRewrite `field:"optional" json:"rewrite" yaml:"rewrite"`
}

