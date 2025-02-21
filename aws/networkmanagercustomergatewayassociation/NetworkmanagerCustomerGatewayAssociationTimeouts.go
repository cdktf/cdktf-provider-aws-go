// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagercustomergatewayassociation


type NetworkmanagerCustomerGatewayAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/networkmanager_customer_gateway_association#create NetworkmanagerCustomerGatewayAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/networkmanager_customer_gateway_association#delete NetworkmanagerCustomerGatewayAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

