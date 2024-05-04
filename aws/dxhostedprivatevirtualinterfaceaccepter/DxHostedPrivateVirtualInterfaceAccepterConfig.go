// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dxhostedprivatevirtualinterfaceaccepter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DxHostedPrivateVirtualInterfaceAccepterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/dx_hosted_private_virtual_interface_accepter#virtual_interface_id DxHostedPrivateVirtualInterfaceAccepter#virtual_interface_id}.
	VirtualInterfaceId *string `field:"required" json:"virtualInterfaceId" yaml:"virtualInterfaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/dx_hosted_private_virtual_interface_accepter#dx_gateway_id DxHostedPrivateVirtualInterfaceAccepter#dx_gateway_id}.
	DxGatewayId *string `field:"optional" json:"dxGatewayId" yaml:"dxGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/dx_hosted_private_virtual_interface_accepter#id DxHostedPrivateVirtualInterfaceAccepter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/dx_hosted_private_virtual_interface_accepter#tags DxHostedPrivateVirtualInterfaceAccepter#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/dx_hosted_private_virtual_interface_accepter#tags_all DxHostedPrivateVirtualInterfaceAccepter#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/dx_hosted_private_virtual_interface_accepter#timeouts DxHostedPrivateVirtualInterfaceAccepter#timeouts}
	Timeouts *DxHostedPrivateVirtualInterfaceAccepterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/dx_hosted_private_virtual_interface_accepter#vpn_gateway_id DxHostedPrivateVirtualInterfaceAccepter#vpn_gateway_id}.
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
}

