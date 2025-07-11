// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3tablestable


type S3TablesTableMaintenanceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/s3tables_table#iceberg_compaction S3TablesTable#iceberg_compaction}.
	IcebergCompaction *S3TablesTableMaintenanceConfigurationIcebergCompaction `field:"optional" json:"icebergCompaction" yaml:"icebergCompaction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/s3tables_table#iceberg_snapshot_management S3TablesTable#iceberg_snapshot_management}.
	IcebergSnapshotManagement *S3TablesTableMaintenanceConfigurationIcebergSnapshotManagement `field:"optional" json:"icebergSnapshotManagement" yaml:"icebergSnapshotManagement"`
}

