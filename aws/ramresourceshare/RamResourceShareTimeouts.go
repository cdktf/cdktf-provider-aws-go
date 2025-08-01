// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ramresourceshare


type RamResourceShareTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ram_resource_share#create RamResourceShare#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ram_resource_share#delete RamResourceShare#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

