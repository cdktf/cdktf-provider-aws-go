// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2clientvpnauthorizationrule


type Ec2ClientVpnAuthorizationRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/ec2_client_vpn_authorization_rule#create Ec2ClientVpnAuthorizationRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/ec2_client_vpn_authorization_rule#delete Ec2ClientVpnAuthorizationRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

