// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivity struct {
	// client_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/msk_cluster#client_authentication MskCluster#client_authentication}
	ClientAuthentication *MskClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthentication `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
}

