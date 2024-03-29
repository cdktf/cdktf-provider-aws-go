// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package finspacekxcluster


type FinspaceKxClusterCapacityConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/finspace_kx_cluster#node_count FinspaceKxCluster#node_count}.
	NodeCount *float64 `field:"required" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/finspace_kx_cluster#node_type FinspaceKxCluster#node_type}.
	NodeType *string `field:"required" json:"nodeType" yaml:"nodeType"`
}

