// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3tablestable


type S3TablesTableMaintenanceConfigurationIcebergCompaction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3tables_table#settings S3TablesTable#settings}.
	Settings *S3TablesTableMaintenanceConfigurationIcebergCompactionSettings `field:"optional" json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3tables_table#status S3TablesTable#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

