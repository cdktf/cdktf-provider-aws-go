// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterBrokerNodeGroupInfoStorageInfo struct {
	// ebs_storage_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/msk_cluster#ebs_storage_info MskCluster#ebs_storage_info}
	EbsStorageInfo *MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfo `field:"optional" json:"ebsStorageInfo" yaml:"ebsStorageInfo"`
}

