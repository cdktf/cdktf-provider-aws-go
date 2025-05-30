// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3tablestablebucket


type S3TablesTableBucketMaintenanceConfigurationIcebergUnreferencedFileRemovalSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/s3tables_table_bucket#non_current_days S3TablesTableBucket#non_current_days}.
	NonCurrentDays *float64 `field:"optional" json:"nonCurrentDays" yaml:"nonCurrentDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/s3tables_table_bucket#unreferenced_days S3TablesTableBucket#unreferenced_days}.
	UnreferencedDays *float64 `field:"optional" json:"unreferencedDays" yaml:"unreferencedDays"`
}

