// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package finspacekxcluster


type FinspaceKxClusterAutoScalingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/finspace_kx_cluster#auto_scaling_metric FinspaceKxCluster#auto_scaling_metric}.
	AutoScalingMetric *string `field:"required" json:"autoScalingMetric" yaml:"autoScalingMetric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/finspace_kx_cluster#max_node_count FinspaceKxCluster#max_node_count}.
	MaxNodeCount *float64 `field:"required" json:"maxNodeCount" yaml:"maxNodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/finspace_kx_cluster#metric_target FinspaceKxCluster#metric_target}.
	MetricTarget *float64 `field:"required" json:"metricTarget" yaml:"metricTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/finspace_kx_cluster#min_node_count FinspaceKxCluster#min_node_count}.
	MinNodeCount *float64 `field:"required" json:"minNodeCount" yaml:"minNodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/finspace_kx_cluster#scale_in_cooldown_seconds FinspaceKxCluster#scale_in_cooldown_seconds}.
	ScaleInCooldownSeconds *float64 `field:"required" json:"scaleInCooldownSeconds" yaml:"scaleInCooldownSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/finspace_kx_cluster#scale_out_cooldown_seconds FinspaceKxCluster#scale_out_cooldown_seconds}.
	ScaleOutCooldownSeconds *float64 `field:"required" json:"scaleOutCooldownSeconds" yaml:"scaleOutCooldownSeconds"`
}

