// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferconnector


type TransferConnectorEgressConfig struct {
	// vpc_lattice block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/transfer_connector#vpc_lattice TransferConnector#vpc_lattice}
	VpcLattice *TransferConnectorEgressConfigVpcLattice `field:"optional" json:"vpcLattice" yaml:"vpcLattice"`
}

