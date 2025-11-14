// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkflowmonitorscope


type NetworkflowmonitorScopeTargetTargetIdentifier struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/networkflowmonitor_scope#target_type NetworkflowmonitorScope#target_type}.
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
	// target_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/networkflowmonitor_scope#target_id NetworkflowmonitorScope#target_id}
	TargetId interface{} `field:"optional" json:"targetId" yaml:"targetId"`
}

