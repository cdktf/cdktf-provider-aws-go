// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskconnectconnector


type MskconnectConnectorKafkaCluster struct {
	// apache_kafka_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/mskconnect_connector#apache_kafka_cluster MskconnectConnector#apache_kafka_cluster}
	ApacheKafkaCluster *MskconnectConnectorKafkaClusterApacheKafkaCluster `field:"required" json:"apacheKafkaCluster" yaml:"apacheKafkaCluster"`
}

