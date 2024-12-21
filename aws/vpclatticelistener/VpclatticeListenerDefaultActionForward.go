// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpclatticelistener


type VpclatticeListenerDefaultActionForward struct {
	// target_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/vpclattice_listener#target_groups VpclatticeListener#target_groups}
	TargetGroups interface{} `field:"optional" json:"targetGroups" yaml:"targetGroups"`
}

