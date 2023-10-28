// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluepartitionindex


type GluePartitionIndexPartitionIndex struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/glue_partition_index#index_name GluePartitionIndex#index_name}.
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/glue_partition_index#keys GluePartitionIndex#keys}.
	Keys *[]*string `field:"optional" json:"keys" yaml:"keys"`
}

