// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistenerrule


type LbListenerRuleConditionHttpHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/lb_listener_rule#http_header_name LbListenerRule#http_header_name}.
	HttpHeaderName *string `field:"required" json:"httpHeaderName" yaml:"httpHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/lb_listener_rule#values LbListenerRule#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

