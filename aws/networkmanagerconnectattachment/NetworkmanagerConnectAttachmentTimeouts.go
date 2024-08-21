// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagerconnectattachment


type NetworkmanagerConnectAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/networkmanager_connect_attachment#create NetworkmanagerConnectAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/networkmanager_connect_attachment#delete NetworkmanagerConnectAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

