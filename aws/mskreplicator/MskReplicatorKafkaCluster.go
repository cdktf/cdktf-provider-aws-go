// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorKafkaCluster struct {
	// amazon_msk_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/msk_replicator#amazon_msk_cluster MskReplicator#amazon_msk_cluster}
	AmazonMskCluster *MskReplicatorKafkaClusterAmazonMskCluster `field:"required" json:"amazonMskCluster" yaml:"amazonMskCluster"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/msk_replicator#vpc_config MskReplicator#vpc_config}
	VpcConfig *MskReplicatorKafkaClusterVpcConfig `field:"required" json:"vpcConfig" yaml:"vpcConfig"`
}

