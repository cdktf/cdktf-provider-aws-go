// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagercustomergatewayassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkmanagerCustomerGatewayAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/networkmanager_customer_gateway_association#customer_gateway_arn NetworkmanagerCustomerGatewayAssociation#customer_gateway_arn}.
	CustomerGatewayArn *string `field:"required" json:"customerGatewayArn" yaml:"customerGatewayArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/networkmanager_customer_gateway_association#device_id NetworkmanagerCustomerGatewayAssociation#device_id}.
	DeviceId *string `field:"required" json:"deviceId" yaml:"deviceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/networkmanager_customer_gateway_association#global_network_id NetworkmanagerCustomerGatewayAssociation#global_network_id}.
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/networkmanager_customer_gateway_association#id NetworkmanagerCustomerGatewayAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/networkmanager_customer_gateway_association#link_id NetworkmanagerCustomerGatewayAssociation#link_id}.
	LinkId *string `field:"optional" json:"linkId" yaml:"linkId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/networkmanager_customer_gateway_association#timeouts NetworkmanagerCustomerGatewayAssociation#timeouts}
	Timeouts *NetworkmanagerCustomerGatewayAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

