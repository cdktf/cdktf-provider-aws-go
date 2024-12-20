// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtableoptimizer


type GlueCatalogTableOptimizerConfigurationOrphanFileDeletionConfigurationIcebergConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/glue_catalog_table_optimizer#location GlueCatalogTableOptimizer#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/glue_catalog_table_optimizer#orphan_file_retention_period_in_days GlueCatalogTableOptimizer#orphan_file_retention_period_in_days}.
	OrphanFileRetentionPeriodInDays *float64 `field:"optional" json:"orphanFileRetentionPeriodInDays" yaml:"orphanFileRetentionPeriodInDays"`
}

