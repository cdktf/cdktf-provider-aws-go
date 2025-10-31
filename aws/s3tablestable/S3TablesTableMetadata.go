// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3tablestable


type S3TablesTableMetadata struct {
	// iceberg block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/s3tables_table#iceberg S3TablesTable#iceberg}
	Iceberg interface{} `field:"optional" json:"iceberg" yaml:"iceberg"`
}

