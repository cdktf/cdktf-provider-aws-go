// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluepartition


type GluePartitionStorageDescriptorSerDeInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/glue_partition#name GluePartition#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/glue_partition#parameters GluePartition#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/glue_partition#serialization_library GluePartition#serialization_library}.
	SerializationLibrary *string `field:"optional" json:"serializationLibrary" yaml:"serializationLibrary"`
}

