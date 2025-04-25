// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtableoptimizer


type GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/glue_catalog_table_optimizer#clean_expired_files GlueCatalogTableOptimizer#clean_expired_files}.
	CleanExpiredFiles interface{} `field:"optional" json:"cleanExpiredFiles" yaml:"cleanExpiredFiles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/glue_catalog_table_optimizer#number_of_snapshots_to_retain GlueCatalogTableOptimizer#number_of_snapshots_to_retain}.
	NumberOfSnapshotsToRetain *float64 `field:"optional" json:"numberOfSnapshotsToRetain" yaml:"numberOfSnapshotsToRetain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/glue_catalog_table_optimizer#snapshot_retention_period_in_days GlueCatalogTableOptimizer#snapshot_retention_period_in_days}.
	SnapshotRetentionPeriodInDays *float64 `field:"optional" json:"snapshotRetentionPeriodInDays" yaml:"snapshotRetentionPeriodInDays"`
}

