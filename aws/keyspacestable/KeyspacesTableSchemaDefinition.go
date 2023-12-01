// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyspacestable


type KeyspacesTableSchemaDefinition struct {
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/keyspaces_table#column KeyspacesTable#column}
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// partition_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/keyspaces_table#partition_key KeyspacesTable#partition_key}
	PartitionKey interface{} `field:"required" json:"partitionKey" yaml:"partitionKey"`
	// clustering_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/keyspaces_table#clustering_key KeyspacesTable#clustering_key}
	ClusteringKey interface{} `field:"optional" json:"clusteringKey" yaml:"clusteringKey"`
	// static_column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/keyspaces_table#static_column KeyspacesTable#static_column}
	StaticColumn interface{} `field:"optional" json:"staticColumn" yaml:"staticColumn"`
}

