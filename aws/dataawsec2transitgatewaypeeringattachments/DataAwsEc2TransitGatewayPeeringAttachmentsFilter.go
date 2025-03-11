// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2transitgatewaypeeringattachments


type DataAwsEc2TransitGatewayPeeringAttachmentsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/ec2_transit_gateway_peering_attachments#name DataAwsEc2TransitGatewayPeeringAttachments#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/ec2_transit_gateway_peering_attachments#values DataAwsEc2TransitGatewayPeeringAttachments#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

