// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2transitgatewayattachment


type DataAwsEc2TransitGatewayAttachmentFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/data-sources/ec2_transit_gateway_attachment#name DataAwsEc2TransitGatewayAttachment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/data-sources/ec2_transit_gateway_attachment#values DataAwsEc2TransitGatewayAttachment#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

