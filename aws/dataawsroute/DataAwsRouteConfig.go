// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsRouteConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/route#route_table_id DataAwsRoute#route_table_id}.
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/route#carrier_gateway_id DataAwsRoute#carrier_gateway_id}.
	CarrierGatewayId *string `field:"optional" json:"carrierGatewayId" yaml:"carrierGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/route#core_network_arn DataAwsRoute#core_network_arn}.
	CoreNetworkArn *string `field:"optional" json:"coreNetworkArn" yaml:"coreNetworkArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/route#destination_cidr_block DataAwsRoute#destination_cidr_block}.
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/route#destination_ipv6_cidr_block DataAwsRoute#destination_ipv6_cidr_block}.
	DestinationIpv6CidrBlock *string `field:"optional" json:"destinationIpv6CidrBlock" yaml:"destinationIpv6CidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/route#destination_prefix_list_id DataAwsRoute#destination_prefix_list_id}.
	DestinationPrefixListId *string `field:"optional" json:"destinationPrefixListId" yaml:"destinationPrefixListId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/route#egress_only_gateway_id DataAwsRoute#egress_only_gateway_id}.
	EgressOnlyGatewayId *string `field:"optional" json:"egressOnlyGatewayId" yaml:"egressOnlyGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/route#gateway_id DataAwsRoute#gateway_id}.
	GatewayId *string `field:"optional" json:"gatewayId" yaml:"gatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/route#id DataAwsRoute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/route#instance_id DataAwsRoute#instance_id}.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/route#local_gateway_id DataAwsRoute#local_gateway_id}.
	LocalGatewayId *string `field:"optional" json:"localGatewayId" yaml:"localGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/route#nat_gateway_id DataAwsRoute#nat_gateway_id}.
	NatGatewayId *string `field:"optional" json:"natGatewayId" yaml:"natGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/route#network_interface_id DataAwsRoute#network_interface_id}.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/route#timeouts DataAwsRoute#timeouts}
	Timeouts *DataAwsRouteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/route#transit_gateway_id DataAwsRoute#transit_gateway_id}.
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/route#vpc_peering_connection_id DataAwsRoute#vpc_peering_connection_id}.
	VpcPeeringConnectionId *string `field:"optional" json:"vpcPeeringConnectionId" yaml:"vpcPeeringConnectionId"`
}

