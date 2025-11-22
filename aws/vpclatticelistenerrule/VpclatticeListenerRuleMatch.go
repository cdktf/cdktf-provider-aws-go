// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpclatticelistenerrule


type VpclatticeListenerRuleMatch struct {
	// http_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/vpclattice_listener_rule#http_match VpclatticeListenerRule#http_match}
	HttpMatch *VpclatticeListenerRuleMatchHttpMatch `field:"required" json:"httpMatch" yaml:"httpMatch"`
}

