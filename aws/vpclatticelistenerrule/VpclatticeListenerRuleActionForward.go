// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpclatticelistenerrule


type VpclatticeListenerRuleActionForward struct {
	// target_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/vpclattice_listener_rule#target_groups VpclatticeListenerRule#target_groups}
	TargetGroups interface{} `field:"required" json:"targetGroups" yaml:"targetGroups"`
}

