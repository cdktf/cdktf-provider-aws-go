// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package neptunecluster


type NeptuneClusterServerlessV2ScalingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/neptune_cluster#max_capacity NeptuneCluster#max_capacity}.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/neptune_cluster#min_capacity NeptuneCluster#min_capacity}.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
}

