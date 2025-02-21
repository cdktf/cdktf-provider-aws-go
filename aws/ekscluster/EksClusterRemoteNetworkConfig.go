// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterRemoteNetworkConfig struct {
	// remote_node_networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/eks_cluster#remote_node_networks EksCluster#remote_node_networks}
	RemoteNodeNetworks *EksClusterRemoteNetworkConfigRemoteNodeNetworks `field:"required" json:"remoteNodeNetworks" yaml:"remoteNodeNetworks"`
	// remote_pod_networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/eks_cluster#remote_pod_networks EksCluster#remote_pod_networks}
	RemotePodNetworks *EksClusterRemoteNetworkConfigRemotePodNetworks `field:"optional" json:"remotePodNetworks" yaml:"remotePodNetworks"`
}

