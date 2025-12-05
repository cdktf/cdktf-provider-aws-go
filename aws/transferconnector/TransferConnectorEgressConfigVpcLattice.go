// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferconnector


type TransferConnectorEgressConfigVpcLattice struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/transfer_connector#resource_configuration_arn TransferConnector#resource_configuration_arn}.
	ResourceConfigurationArn *string `field:"required" json:"resourceConfigurationArn" yaml:"resourceConfigurationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/transfer_connector#port_number TransferConnector#port_number}.
	PortNumber *float64 `field:"optional" json:"portNumber" yaml:"portNumber"`
}

