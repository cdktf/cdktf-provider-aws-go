// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskconnectconnector


type MskconnectConnectorCapacityProvisionedCapacity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/mskconnect_connector#worker_count MskconnectConnector#worker_count}.
	WorkerCount *float64 `field:"required" json:"workerCount" yaml:"workerCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/mskconnect_connector#mcu_count MskconnectConnector#mcu_count}.
	McuCount *float64 `field:"optional" json:"mcuCount" yaml:"mcuCount"`
}

