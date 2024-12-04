// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/msk_cluster#enabled MskCluster#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/msk_cluster#volume_throughput MskCluster#volume_throughput}.
	VolumeThroughput *float64 `field:"optional" json:"volumeThroughput" yaml:"volumeThroughput"`
}

