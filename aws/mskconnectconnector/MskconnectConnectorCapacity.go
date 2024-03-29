// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskconnectconnector


type MskconnectConnectorCapacity struct {
	// autoscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/mskconnect_connector#autoscaling MskconnectConnector#autoscaling}
	Autoscaling *MskconnectConnectorCapacityAutoscaling `field:"optional" json:"autoscaling" yaml:"autoscaling"`
	// provisioned_capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/mskconnect_connector#provisioned_capacity MskconnectConnector#provisioned_capacity}
	ProvisionedCapacity *MskconnectConnectorCapacityProvisionedCapacity `field:"optional" json:"provisionedCapacity" yaml:"provisionedCapacity"`
}

