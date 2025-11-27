// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dxprivatevirtualinterface

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DxPrivateVirtualInterfaceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/dx_private_virtual_interface#address_family DxPrivateVirtualInterface#address_family}.
	AddressFamily *string `field:"required" json:"addressFamily" yaml:"addressFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/dx_private_virtual_interface#bgp_asn DxPrivateVirtualInterface#bgp_asn}.
	BgpAsn *float64 `field:"required" json:"bgpAsn" yaml:"bgpAsn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/dx_private_virtual_interface#connection_id DxPrivateVirtualInterface#connection_id}.
	ConnectionId *string `field:"required" json:"connectionId" yaml:"connectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/dx_private_virtual_interface#name DxPrivateVirtualInterface#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/dx_private_virtual_interface#vlan DxPrivateVirtualInterface#vlan}.
	Vlan *float64 `field:"required" json:"vlan" yaml:"vlan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/dx_private_virtual_interface#amazon_address DxPrivateVirtualInterface#amazon_address}.
	AmazonAddress *string `field:"optional" json:"amazonAddress" yaml:"amazonAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/dx_private_virtual_interface#bgp_auth_key DxPrivateVirtualInterface#bgp_auth_key}.
	BgpAuthKey *string `field:"optional" json:"bgpAuthKey" yaml:"bgpAuthKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/dx_private_virtual_interface#customer_address DxPrivateVirtualInterface#customer_address}.
	CustomerAddress *string `field:"optional" json:"customerAddress" yaml:"customerAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/dx_private_virtual_interface#dx_gateway_id DxPrivateVirtualInterface#dx_gateway_id}.
	DxGatewayId *string `field:"optional" json:"dxGatewayId" yaml:"dxGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/dx_private_virtual_interface#id DxPrivateVirtualInterface#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/dx_private_virtual_interface#mtu DxPrivateVirtualInterface#mtu}.
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/dx_private_virtual_interface#region DxPrivateVirtualInterface#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/dx_private_virtual_interface#sitelink_enabled DxPrivateVirtualInterface#sitelink_enabled}.
	SitelinkEnabled interface{} `field:"optional" json:"sitelinkEnabled" yaml:"sitelinkEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/dx_private_virtual_interface#tags DxPrivateVirtualInterface#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/dx_private_virtual_interface#tags_all DxPrivateVirtualInterface#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/dx_private_virtual_interface#timeouts DxPrivateVirtualInterface#timeouts}
	Timeouts *DxPrivateVirtualInterfaceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/dx_private_virtual_interface#vpn_gateway_id DxPrivateVirtualInterface#vpn_gateway_id}.
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
}

