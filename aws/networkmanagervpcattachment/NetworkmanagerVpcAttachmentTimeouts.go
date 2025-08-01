// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagervpcattachment


type NetworkmanagerVpcAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkmanager_vpc_attachment#create NetworkmanagerVpcAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkmanager_vpc_attachment#delete NetworkmanagerVpcAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkmanager_vpc_attachment#update NetworkmanagerVpcAttachment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

