// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpclatticelistenerrule


type VpclatticeListenerRuleMatchHttpMatch struct {
	// header_matches block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/vpclattice_listener_rule#header_matches VpclatticeListenerRule#header_matches}
	HeaderMatches interface{} `field:"optional" json:"headerMatches" yaml:"headerMatches"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/vpclattice_listener_rule#method VpclatticeListenerRule#method}.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// path_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/vpclattice_listener_rule#path_match VpclatticeListenerRule#path_match}
	PathMatch *VpclatticeListenerRuleMatchHttpMatchPathMatch `field:"optional" json:"pathMatch" yaml:"pathMatch"`
}

