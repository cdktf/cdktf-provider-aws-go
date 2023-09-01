// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2transitgatewayroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2TransitGatewayRouteConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/ec2_transit_gateway_route#destination_cidr_block Ec2TransitGatewayRoute#destination_cidr_block}.
	DestinationCidrBlock *string `field:"required" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/ec2_transit_gateway_route#transit_gateway_route_table_id Ec2TransitGatewayRoute#transit_gateway_route_table_id}.
	TransitGatewayRouteTableId *string `field:"required" json:"transitGatewayRouteTableId" yaml:"transitGatewayRouteTableId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/ec2_transit_gateway_route#blackhole Ec2TransitGatewayRoute#blackhole}.
	Blackhole interface{} `field:"optional" json:"blackhole" yaml:"blackhole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/ec2_transit_gateway_route#id Ec2TransitGatewayRoute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/ec2_transit_gateway_route#transit_gateway_attachment_id Ec2TransitGatewayRoute#transit_gateway_attachment_id}.
	TransitGatewayAttachmentId *string `field:"optional" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
}

