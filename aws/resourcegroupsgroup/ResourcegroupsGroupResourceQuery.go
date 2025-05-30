// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcegroupsgroup


type ResourcegroupsGroupResourceQuery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/resourcegroups_group#query ResourcegroupsGroup#query}.
	Query *string `field:"required" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/resourcegroups_group#type ResourcegroupsGroup#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

