// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2managedprefixlists


type DataAwsEc2ManagedPrefixListsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/data-sources/ec2_managed_prefix_lists#name DataAwsEc2ManagedPrefixLists#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/data-sources/ec2_managed_prefix_lists#values DataAwsEc2ManagedPrefixLists#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

