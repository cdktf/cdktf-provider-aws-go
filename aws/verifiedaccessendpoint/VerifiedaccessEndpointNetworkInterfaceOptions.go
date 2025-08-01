// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedaccessendpoint


type VerifiedaccessEndpointNetworkInterfaceOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/verifiedaccess_endpoint#network_interface_id VerifiedaccessEndpoint#network_interface_id}.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/verifiedaccess_endpoint#port VerifiedaccessEndpoint#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/verifiedaccess_endpoint#port_range VerifiedaccessEndpoint#port_range}
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/verifiedaccess_endpoint#protocol VerifiedaccessEndpoint#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

