// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eksnodegroup


type EksNodeGroupUpdateConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/eks_node_group#max_unavailable EksNodeGroup#max_unavailable}.
	MaxUnavailable *float64 `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/eks_node_group#max_unavailable_percentage EksNodeGroup#max_unavailable_percentage}.
	MaxUnavailablePercentage *float64 `field:"optional" json:"maxUnavailablePercentage" yaml:"maxUnavailablePercentage"`
}

