// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterOpenMonitoring struct {
	// prometheus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/msk_cluster#prometheus MskCluster#prometheus}
	Prometheus *MskClusterOpenMonitoringPrometheus `field:"required" json:"prometheus" yaml:"prometheus"`
}

