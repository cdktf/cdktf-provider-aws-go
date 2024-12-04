// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawslblistenerrule


type DataAwsLbListenerRuleCondition struct {
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/data-sources/lb_listener_rule#query_string DataAwsLbListenerRule#query_string}
	QueryString *DataAwsLbListenerRuleConditionQueryString `field:"optional" json:"queryString" yaml:"queryString"`
}

