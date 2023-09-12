// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alblistenerrule


type AlbListenerRuleConditionHttpRequestMethod struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/alb_listener_rule#values AlbListenerRule#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

