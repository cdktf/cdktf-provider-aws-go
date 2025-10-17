// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3tablestablebucket


type S3TablesTableBucketMaintenanceConfigurationIcebergUnreferencedFileRemoval struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/s3tables_table_bucket#settings S3TablesTableBucket#settings}.
	Settings *S3TablesTableBucketMaintenanceConfigurationIcebergUnreferencedFileRemovalSettings `field:"optional" json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/s3tables_table_bucket#status S3TablesTableBucket#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

