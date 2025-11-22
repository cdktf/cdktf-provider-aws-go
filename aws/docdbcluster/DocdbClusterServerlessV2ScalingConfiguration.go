// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package docdbcluster


type DocdbClusterServerlessV2ScalingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/docdb_cluster#max_capacity DocdbCluster#max_capacity}.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/docdb_cluster#min_capacity DocdbCluster#min_capacity}.
	MinCapacity *float64 `field:"required" json:"minCapacity" yaml:"minCapacity"`
}

