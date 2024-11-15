// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcpeeringconnectionaccepter


type VpcPeeringConnectionAccepterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/vpc_peering_connection_accepter#create VpcPeeringConnectionAccepterA#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/vpc_peering_connection_accepter#update VpcPeeringConnectionAccepterA#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

