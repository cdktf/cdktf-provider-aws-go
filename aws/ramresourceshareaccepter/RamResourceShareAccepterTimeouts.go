// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ramresourceshareaccepter


type RamResourceShareAccepterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ram_resource_share_accepter#create RamResourceShareAccepter#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ram_resource_share_accepter#delete RamResourceShareAccepter#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

