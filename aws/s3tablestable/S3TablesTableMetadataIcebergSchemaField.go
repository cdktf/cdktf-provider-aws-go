// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3tablestable


type S3TablesTableMetadataIcebergSchemaField struct {
	// The name of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/s3tables_table#name S3TablesTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The field type. S3 Tables supports all Apache Iceberg primitive types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/s3tables_table#type S3TablesTable#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A Boolean value that specifies whether values are required for each row in this field. Default: false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/s3tables_table#required S3TablesTable#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

