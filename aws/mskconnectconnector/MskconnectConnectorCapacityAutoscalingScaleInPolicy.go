// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskconnectconnector


type MskconnectConnectorCapacityAutoscalingScaleInPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/mskconnect_connector#cpu_utilization_percentage MskconnectConnector#cpu_utilization_percentage}.
	CpuUtilizationPercentage *float64 `field:"optional" json:"cpuUtilizationPercentage" yaml:"cpuUtilizationPercentage"`
}

