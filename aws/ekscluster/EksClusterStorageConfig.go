// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterStorageConfig struct {
	// block_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/eks_cluster#block_storage EksCluster#block_storage}
	BlockStorage *EksClusterStorageConfigBlockStorage `field:"optional" json:"blockStorage" yaml:"blockStorage"`
}

