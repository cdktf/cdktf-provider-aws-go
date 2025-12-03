// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawslblistenerrule


type DataAwsLbListenerRuleActionForward struct {
	// stickiness block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/data-sources/lb_listener_rule#stickiness DataAwsLbListenerRule#stickiness}
	Stickiness interface{} `field:"optional" json:"stickiness" yaml:"stickiness"`
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/data-sources/lb_listener_rule#target_group DataAwsLbListenerRule#target_group}
	TargetGroup interface{} `field:"optional" json:"targetGroup" yaml:"targetGroup"`
}

