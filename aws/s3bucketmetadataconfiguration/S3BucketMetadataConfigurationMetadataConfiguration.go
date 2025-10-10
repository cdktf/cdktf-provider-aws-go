// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketmetadataconfiguration


type S3BucketMetadataConfigurationMetadataConfiguration struct {
	// inventory_table_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/s3_bucket_metadata_configuration#inventory_table_configuration S3BucketMetadataConfiguration#inventory_table_configuration}
	InventoryTableConfiguration interface{} `field:"optional" json:"inventoryTableConfiguration" yaml:"inventoryTableConfiguration"`
	// journal_table_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/s3_bucket_metadata_configuration#journal_table_configuration S3BucketMetadataConfiguration#journal_table_configuration}
	JournalTableConfiguration interface{} `field:"optional" json:"journalTableConfiguration" yaml:"journalTableConfiguration"`
}

