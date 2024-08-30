// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpclatticelistenerrule


type VpclatticeListenerRuleAction struct {
	// fixed_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/vpclattice_listener_rule#fixed_response VpclatticeListenerRule#fixed_response}
	FixedResponse *VpclatticeListenerRuleActionFixedResponse `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// forward block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/vpclattice_listener_rule#forward VpclatticeListenerRule#forward}
	Forward *VpclatticeListenerRuleActionForward `field:"optional" json:"forward" yaml:"forward"`
}

