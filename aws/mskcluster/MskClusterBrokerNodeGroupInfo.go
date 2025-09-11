// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterBrokerNodeGroupInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/msk_cluster#client_subnets MskCluster#client_subnets}.
	ClientSubnets *[]*string `field:"required" json:"clientSubnets" yaml:"clientSubnets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/msk_cluster#instance_type MskCluster#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/msk_cluster#security_groups MskCluster#security_groups}.
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/msk_cluster#az_distribution MskCluster#az_distribution}.
	AzDistribution *string `field:"optional" json:"azDistribution" yaml:"azDistribution"`
	// connectivity_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/msk_cluster#connectivity_info MskCluster#connectivity_info}
	ConnectivityInfo *MskClusterBrokerNodeGroupInfoConnectivityInfo `field:"optional" json:"connectivityInfo" yaml:"connectivityInfo"`
	// storage_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/msk_cluster#storage_info MskCluster#storage_info}
	StorageInfo *MskClusterBrokerNodeGroupInfoStorageInfo `field:"optional" json:"storageInfo" yaml:"storageInfo"`
}

