// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable


type GlueCatalogTableStorageDescriptor struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/glue_catalog_table#bucket_columns GlueCatalogTable#bucket_columns}.
	BucketColumns *[]*string `field:"optional" json:"bucketColumns" yaml:"bucketColumns"`
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/glue_catalog_table#columns GlueCatalogTable#columns}
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/glue_catalog_table#compressed GlueCatalogTable#compressed}.
	Compressed interface{} `field:"optional" json:"compressed" yaml:"compressed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/glue_catalog_table#input_format GlueCatalogTable#input_format}.
	InputFormat *string `field:"optional" json:"inputFormat" yaml:"inputFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/glue_catalog_table#location GlueCatalogTable#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/glue_catalog_table#number_of_buckets GlueCatalogTable#number_of_buckets}.
	NumberOfBuckets *float64 `field:"optional" json:"numberOfBuckets" yaml:"numberOfBuckets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/glue_catalog_table#output_format GlueCatalogTable#output_format}.
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/glue_catalog_table#parameters GlueCatalogTable#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// schema_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/glue_catalog_table#schema_reference GlueCatalogTable#schema_reference}
	SchemaReference *GlueCatalogTableStorageDescriptorSchemaReference `field:"optional" json:"schemaReference" yaml:"schemaReference"`
	// ser_de_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/glue_catalog_table#ser_de_info GlueCatalogTable#ser_de_info}
	SerDeInfo *GlueCatalogTableStorageDescriptorSerDeInfo `field:"optional" json:"serDeInfo" yaml:"serDeInfo"`
	// skewed_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/glue_catalog_table#skewed_info GlueCatalogTable#skewed_info}
	SkewedInfo *GlueCatalogTableStorageDescriptorSkewedInfo `field:"optional" json:"skewedInfo" yaml:"skewedInfo"`
	// sort_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/glue_catalog_table#sort_columns GlueCatalogTable#sort_columns}
	SortColumns interface{} `field:"optional" json:"sortColumns" yaml:"sortColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/glue_catalog_table#stored_as_sub_directories GlueCatalogTable#stored_as_sub_directories}.
	StoredAsSubDirectories interface{} `field:"optional" json:"storedAsSubDirectories" yaml:"storedAsSubDirectories"`
}

