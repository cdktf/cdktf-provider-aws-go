// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluepartition


type GluePartitionStorageDescriptor struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/glue_partition#bucket_columns GluePartition#bucket_columns}.
	BucketColumns *[]*string `field:"optional" json:"bucketColumns" yaml:"bucketColumns"`
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/glue_partition#columns GluePartition#columns}
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/glue_partition#compressed GluePartition#compressed}.
	Compressed interface{} `field:"optional" json:"compressed" yaml:"compressed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/glue_partition#input_format GluePartition#input_format}.
	InputFormat *string `field:"optional" json:"inputFormat" yaml:"inputFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/glue_partition#location GluePartition#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/glue_partition#number_of_buckets GluePartition#number_of_buckets}.
	NumberOfBuckets *float64 `field:"optional" json:"numberOfBuckets" yaml:"numberOfBuckets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/glue_partition#output_format GluePartition#output_format}.
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/glue_partition#parameters GluePartition#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// ser_de_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/glue_partition#ser_de_info GluePartition#ser_de_info}
	SerDeInfo *GluePartitionStorageDescriptorSerDeInfo `field:"optional" json:"serDeInfo" yaml:"serDeInfo"`
	// skewed_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/glue_partition#skewed_info GluePartition#skewed_info}
	SkewedInfo *GluePartitionStorageDescriptorSkewedInfo `field:"optional" json:"skewedInfo" yaml:"skewedInfo"`
	// sort_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/glue_partition#sort_columns GluePartition#sort_columns}
	SortColumns interface{} `field:"optional" json:"sortColumns" yaml:"sortColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/glue_partition#stored_as_sub_directories GluePartition#stored_as_sub_directories}.
	StoredAsSubDirectories interface{} `field:"optional" json:"storedAsSubDirectories" yaml:"storedAsSubDirectories"`
}

