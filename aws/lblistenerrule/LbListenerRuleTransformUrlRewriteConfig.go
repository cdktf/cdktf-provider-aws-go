// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistenerrule


type LbListenerRuleTransformUrlRewriteConfig struct {
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/lb_listener_rule#rewrite LbListenerRule#rewrite}
	Rewrite *LbListenerRuleTransformUrlRewriteConfigRewrite `field:"optional" json:"rewrite" yaml:"rewrite"`
}

