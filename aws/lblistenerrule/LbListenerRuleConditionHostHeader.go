// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistenerrule


type LbListenerRuleConditionHostHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/lb_listener_rule#values LbListenerRule#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

