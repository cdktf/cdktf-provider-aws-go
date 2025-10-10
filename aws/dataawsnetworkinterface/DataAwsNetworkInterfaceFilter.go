// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsnetworkinterface


type DataAwsNetworkInterfaceFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/data-sources/network_interface#name DataAwsNetworkInterface#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/data-sources/network_interface#values DataAwsNetworkInterface#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

