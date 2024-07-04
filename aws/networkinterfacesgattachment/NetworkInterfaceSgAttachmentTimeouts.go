// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkinterfacesgattachment


type NetworkInterfaceSgAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/network_interface_sg_attachment#create NetworkInterfaceSgAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/network_interface_sg_attachment#delete NetworkInterfaceSgAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/network_interface_sg_attachment#read NetworkInterfaceSgAttachment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

