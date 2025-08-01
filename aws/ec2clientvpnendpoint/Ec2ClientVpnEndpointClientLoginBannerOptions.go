// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2clientvpnendpoint


type Ec2ClientVpnEndpointClientLoginBannerOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ec2_client_vpn_endpoint#banner_text Ec2ClientVpnEndpoint#banner_text}.
	BannerText *string `field:"optional" json:"bannerText" yaml:"bannerText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ec2_client_vpn_endpoint#enabled Ec2ClientVpnEndpoint#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

