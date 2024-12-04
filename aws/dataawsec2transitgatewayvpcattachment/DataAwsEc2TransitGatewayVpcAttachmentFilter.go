// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2transitgatewayvpcattachment


type DataAwsEc2TransitGatewayVpcAttachmentFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/data-sources/ec2_transit_gateway_vpc_attachment#name DataAwsEc2TransitGatewayVpcAttachment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/data-sources/ec2_transit_gateway_vpc_attachment#values DataAwsEc2TransitGatewayVpcAttachment#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

