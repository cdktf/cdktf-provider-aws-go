// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawslblistenerrule


type DataAwsLbListenerRuleAction struct {
	// forward block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/data-sources/lb_listener_rule#forward DataAwsLbListenerRule#forward}
	Forward *DataAwsLbListenerRuleActionForward `field:"optional" json:"forward" yaml:"forward"`
}

