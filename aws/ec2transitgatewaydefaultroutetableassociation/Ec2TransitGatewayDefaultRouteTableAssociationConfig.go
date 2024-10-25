// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2transitgatewaydefaultroutetableassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2TransitGatewayDefaultRouteTableAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/ec2_transit_gateway_default_route_table_association#transit_gateway_id Ec2TransitGatewayDefaultRouteTableAssociation#transit_gateway_id}.
	TransitGatewayId *string `field:"required" json:"transitGatewayId" yaml:"transitGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/ec2_transit_gateway_default_route_table_association#transit_gateway_route_table_id Ec2TransitGatewayDefaultRouteTableAssociation#transit_gateway_route_table_id}.
	TransitGatewayRouteTableId *string `field:"required" json:"transitGatewayRouteTableId" yaml:"transitGatewayRouteTableId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/ec2_transit_gateway_default_route_table_association#timeouts Ec2TransitGatewayDefaultRouteTableAssociation#timeouts}
	Timeouts *Ec2TransitGatewayDefaultRouteTableAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

