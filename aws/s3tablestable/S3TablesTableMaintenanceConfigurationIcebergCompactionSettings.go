// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3tablestable


type S3TablesTableMaintenanceConfigurationIcebergCompactionSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.1/docs/resources/s3tables_table#target_file_size_mb S3TablesTable#target_file_size_mb}.
	TargetFileSizeMb *float64 `field:"optional" json:"targetFileSizeMb" yaml:"targetFileSizeMb"`
}

