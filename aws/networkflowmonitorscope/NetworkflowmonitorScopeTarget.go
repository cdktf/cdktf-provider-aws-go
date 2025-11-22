// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkflowmonitorscope


type NetworkflowmonitorScopeTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/networkflowmonitor_scope#region NetworkflowmonitorScope#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// target_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/networkflowmonitor_scope#target_identifier NetworkflowmonitorScope#target_identifier}
	TargetIdentifier interface{} `field:"optional" json:"targetIdentifier" yaml:"targetIdentifier"`
}

