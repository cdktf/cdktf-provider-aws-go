// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfo struct {
	// provisioned_throughput block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/msk_cluster#provisioned_throughput MskCluster#provisioned_throughput}
	ProvisionedThroughput *MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughput `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/msk_cluster#volume_size MskCluster#volume_size}.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
}

