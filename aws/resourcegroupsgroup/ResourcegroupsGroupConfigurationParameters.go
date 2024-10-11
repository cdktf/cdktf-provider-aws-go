// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcegroupsgroup


type ResourcegroupsGroupConfigurationParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/resourcegroups_group#name ResourcegroupsGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/resourcegroups_group#values ResourcegroupsGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

