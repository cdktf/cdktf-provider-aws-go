// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2transitgatewaydxgatewayattachment


type DataAwsEc2TransitGatewayDxGatewayAttachmentFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/data-sources/ec2_transit_gateway_dx_gateway_attachment#name DataAwsEc2TransitGatewayDxGatewayAttachment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/data-sources/ec2_transit_gateway_dx_gateway_attachment#values DataAwsEc2TransitGatewayDxGatewayAttachment#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

