// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterOpenMonitoringPrometheus struct {
	// jmx_exporter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/msk_cluster#jmx_exporter MskCluster#jmx_exporter}
	JmxExporter *MskClusterOpenMonitoringPrometheusJmxExporter `field:"optional" json:"jmxExporter" yaml:"jmxExporter"`
	// node_exporter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/msk_cluster#node_exporter MskCluster#node_exporter}
	NodeExporter *MskClusterOpenMonitoringPrometheusNodeExporter `field:"optional" json:"nodeExporter" yaml:"nodeExporter"`
}

