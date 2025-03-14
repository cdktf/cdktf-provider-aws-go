// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagertransitgatewayroutetableattachment


type NetworkmanagerTransitGatewayRouteTableAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/networkmanager_transit_gateway_route_table_attachment#create NetworkmanagerTransitGatewayRouteTableAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/networkmanager_transit_gateway_route_table_attachment#delete NetworkmanagerTransitGatewayRouteTableAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

