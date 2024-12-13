// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eksnodegroup


type EksNodeGroupScalingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/eks_node_group#desired_size EksNodeGroup#desired_size}.
	DesiredSize *float64 `field:"required" json:"desiredSize" yaml:"desiredSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/eks_node_group#max_size EksNodeGroup#max_size}.
	MaxSize *float64 `field:"required" json:"maxSize" yaml:"maxSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/eks_node_group#min_size EksNodeGroup#min_size}.
	MinSize *float64 `field:"required" json:"minSize" yaml:"minSize"`
}

