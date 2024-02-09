// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterBrokerNodeGroupInfoConnectivityInfo struct {
	// public_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/msk_cluster#public_access MskCluster#public_access}
	PublicAccess *MskClusterBrokerNodeGroupInfoConnectivityInfoPublicAccess `field:"optional" json:"publicAccess" yaml:"publicAccess"`
	// vpc_connectivity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/msk_cluster#vpc_connectivity MskCluster#vpc_connectivity}
	VpcConnectivity *MskClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivity `field:"optional" json:"vpcConnectivity" yaml:"vpcConnectivity"`
}

