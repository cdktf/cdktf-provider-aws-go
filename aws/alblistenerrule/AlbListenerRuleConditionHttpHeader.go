// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alblistenerrule


type AlbListenerRuleConditionHttpHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/alb_listener_rule#http_header_name AlbListenerRule#http_header_name}.
	HttpHeaderName *string `field:"required" json:"httpHeaderName" yaml:"httpHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/alb_listener_rule#regex_values AlbListenerRule#regex_values}.
	RegexValues *[]*string `field:"optional" json:"regexValues" yaml:"regexValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/alb_listener_rule#values AlbListenerRule#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

