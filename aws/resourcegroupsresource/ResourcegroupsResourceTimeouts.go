// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcegroupsresource


type ResourcegroupsResourceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/resourcegroups_resource#create ResourcegroupsResource#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/resourcegroups_resource#delete ResourcegroupsResource#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

