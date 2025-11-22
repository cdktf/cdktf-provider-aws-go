// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterRemoteNetworkConfigRemoteNodeNetworks struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/eks_cluster#cidrs EksCluster#cidrs}.
	Cidrs *[]*string `field:"optional" json:"cidrs" yaml:"cidrs"`
}

