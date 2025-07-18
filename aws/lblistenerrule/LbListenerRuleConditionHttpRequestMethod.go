// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistenerrule


type LbListenerRuleConditionHttpRequestMethod struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/lb_listener_rule#values LbListenerRule#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

