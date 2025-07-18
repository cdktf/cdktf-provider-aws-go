// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcegroupsgroup


type ResourcegroupsGroupConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/resourcegroups_group#type ResourcegroupsGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/resourcegroups_group#parameters ResourcegroupsGroup#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

