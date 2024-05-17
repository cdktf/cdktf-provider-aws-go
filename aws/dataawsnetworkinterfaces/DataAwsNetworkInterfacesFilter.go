// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsnetworkinterfaces


type DataAwsNetworkInterfacesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.50.0/docs/data-sources/network_interfaces#name DataAwsNetworkInterfaces#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.50.0/docs/data-sources/network_interfaces#values DataAwsNetworkInterfaces#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

