// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2transitgatewayconnectpeer


type Ec2TransitGatewayConnectPeerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/ec2_transit_gateway_connect_peer#create Ec2TransitGatewayConnectPeer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/ec2_transit_gateway_connect_peer#delete Ec2TransitGatewayConnectPeer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

