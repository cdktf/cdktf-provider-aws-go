// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpclatticelistenerrule


type VpclatticeListenerRuleMatchHttpMatchPathMatchMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/vpclattice_listener_rule#exact VpclatticeListenerRule#exact}.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/vpclattice_listener_rule#prefix VpclatticeListenerRule#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

