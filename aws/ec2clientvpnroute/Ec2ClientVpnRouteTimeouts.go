// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2clientvpnroute


type Ec2ClientVpnRouteTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/ec2_client_vpn_route#create Ec2ClientVpnRoute#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/ec2_client_vpn_route#delete Ec2ClientVpnRoute#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

