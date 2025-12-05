// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3tablestable


type S3TablesTableMetadataIceberg struct {
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/s3tables_table#schema S3TablesTable#schema}
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
}

