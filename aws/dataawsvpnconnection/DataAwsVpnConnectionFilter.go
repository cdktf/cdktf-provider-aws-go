// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsvpnconnection


type DataAwsVpnConnectionFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/data-sources/vpn_connection#name DataAwsVpnConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/data-sources/vpn_connection#values DataAwsVpnConnection#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

