// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagertransitgatewayregistration


type NetworkmanagerTransitGatewayRegistrationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/networkmanager_transit_gateway_registration#create NetworkmanagerTransitGatewayRegistration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/networkmanager_transit_gateway_registration#delete NetworkmanagerTransitGatewayRegistration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

