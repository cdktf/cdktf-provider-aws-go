// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alblistenerrule


type AlbListenerRuleConditionHostHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/alb_listener_rule#regex_values AlbListenerRule#regex_values}.
	RegexValues *[]*string `field:"optional" json:"regexValues" yaml:"regexValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/alb_listener_rule#values AlbListenerRule#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

