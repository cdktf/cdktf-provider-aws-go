// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2clientvpnauthorizationrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2ClientVpnAuthorizationRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/ec2_client_vpn_authorization_rule#client_vpn_endpoint_id Ec2ClientVpnAuthorizationRule#client_vpn_endpoint_id}.
	ClientVpnEndpointId *string `field:"required" json:"clientVpnEndpointId" yaml:"clientVpnEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/ec2_client_vpn_authorization_rule#target_network_cidr Ec2ClientVpnAuthorizationRule#target_network_cidr}.
	TargetNetworkCidr *string `field:"required" json:"targetNetworkCidr" yaml:"targetNetworkCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/ec2_client_vpn_authorization_rule#access_group_id Ec2ClientVpnAuthorizationRule#access_group_id}.
	AccessGroupId *string `field:"optional" json:"accessGroupId" yaml:"accessGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/ec2_client_vpn_authorization_rule#authorize_all_groups Ec2ClientVpnAuthorizationRule#authorize_all_groups}.
	AuthorizeAllGroups interface{} `field:"optional" json:"authorizeAllGroups" yaml:"authorizeAllGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/ec2_client_vpn_authorization_rule#description Ec2ClientVpnAuthorizationRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/ec2_client_vpn_authorization_rule#id Ec2ClientVpnAuthorizationRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/ec2_client_vpn_authorization_rule#timeouts Ec2ClientVpnAuthorizationRule#timeouts}
	Timeouts *Ec2ClientVpnAuthorizationRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

