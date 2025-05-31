// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package finspacekxcluster


type FinspaceKxClusterScalingGroupConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/finspace_kx_cluster#memory_reservation FinspaceKxCluster#memory_reservation}.
	MemoryReservation *float64 `field:"required" json:"memoryReservation" yaml:"memoryReservation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/finspace_kx_cluster#node_count FinspaceKxCluster#node_count}.
	NodeCount *float64 `field:"required" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/finspace_kx_cluster#scaling_group_name FinspaceKxCluster#scaling_group_name}.
	ScalingGroupName *string `field:"required" json:"scalingGroupName" yaml:"scalingGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/finspace_kx_cluster#cpu FinspaceKxCluster#cpu}.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/finspace_kx_cluster#memory_limit FinspaceKxCluster#memory_limit}.
	MemoryLimit *float64 `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
}

