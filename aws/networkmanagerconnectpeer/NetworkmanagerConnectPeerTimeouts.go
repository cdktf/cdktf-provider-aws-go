// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagerconnectpeer


type NetworkmanagerConnectPeerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/networkmanager_connect_peer#create NetworkmanagerConnectPeer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/networkmanager_connect_peer#delete NetworkmanagerConnectPeer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

