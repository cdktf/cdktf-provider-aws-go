// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3tablestablebucket


type S3TablesTableBucketMaintenanceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/s3tables_table_bucket#iceberg_unreferenced_file_removal S3TablesTableBucket#iceberg_unreferenced_file_removal}.
	IcebergUnreferencedFileRemoval *S3TablesTableBucketMaintenanceConfigurationIcebergUnreferencedFileRemoval `field:"optional" json:"icebergUnreferencedFileRemoval" yaml:"icebergUnreferencedFileRemoval"`
}

