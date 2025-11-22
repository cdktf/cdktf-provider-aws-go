// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eksnodegroup


type EksNodeGroupNodeRepairConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/eks_node_group#enabled EksNodeGroup#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/eks_node_group#max_parallel_nodes_repaired_count EksNodeGroup#max_parallel_nodes_repaired_count}.
	MaxParallelNodesRepairedCount *float64 `field:"optional" json:"maxParallelNodesRepairedCount" yaml:"maxParallelNodesRepairedCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/eks_node_group#max_parallel_nodes_repaired_percentage EksNodeGroup#max_parallel_nodes_repaired_percentage}.
	MaxParallelNodesRepairedPercentage *float64 `field:"optional" json:"maxParallelNodesRepairedPercentage" yaml:"maxParallelNodesRepairedPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/eks_node_group#max_unhealthy_node_threshold_count EksNodeGroup#max_unhealthy_node_threshold_count}.
	MaxUnhealthyNodeThresholdCount *float64 `field:"optional" json:"maxUnhealthyNodeThresholdCount" yaml:"maxUnhealthyNodeThresholdCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/eks_node_group#max_unhealthy_node_threshold_percentage EksNodeGroup#max_unhealthy_node_threshold_percentage}.
	MaxUnhealthyNodeThresholdPercentage *float64 `field:"optional" json:"maxUnhealthyNodeThresholdPercentage" yaml:"maxUnhealthyNodeThresholdPercentage"`
	// node_repair_config_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/eks_node_group#node_repair_config_overrides EksNodeGroup#node_repair_config_overrides}
	NodeRepairConfigOverrides interface{} `field:"optional" json:"nodeRepairConfigOverrides" yaml:"nodeRepairConfigOverrides"`
}

