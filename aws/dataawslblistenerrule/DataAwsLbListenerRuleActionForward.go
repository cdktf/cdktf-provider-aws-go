// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawslblistenerrule


type DataAwsLbListenerRuleActionForward struct {
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/data-sources/lb_listener_rule#target_group DataAwsLbListenerRule#target_group}
	TargetGroup interface{} `field:"optional" json:"targetGroup" yaml:"targetGroup"`
}

