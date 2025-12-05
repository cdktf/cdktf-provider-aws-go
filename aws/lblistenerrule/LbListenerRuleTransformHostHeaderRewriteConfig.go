// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistenerrule


type LbListenerRuleTransformHostHeaderRewriteConfig struct {
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lb_listener_rule#rewrite LbListenerRule#rewrite}
	Rewrite *LbListenerRuleTransformHostHeaderRewriteConfigRewrite `field:"optional" json:"rewrite" yaml:"rewrite"`
}

