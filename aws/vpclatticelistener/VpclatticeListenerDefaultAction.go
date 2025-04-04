// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpclatticelistener


type VpclatticeListenerDefaultAction struct {
	// fixed_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/vpclattice_listener#fixed_response VpclatticeListener#fixed_response}
	FixedResponse *VpclatticeListenerDefaultActionFixedResponse `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// forward block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/vpclattice_listener#forward VpclatticeListener#forward}
	Forward interface{} `field:"optional" json:"forward" yaml:"forward"`
}

