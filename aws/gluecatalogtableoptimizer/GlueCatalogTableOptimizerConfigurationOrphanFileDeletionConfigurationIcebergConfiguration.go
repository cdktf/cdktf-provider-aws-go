// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtableoptimizer


type GlueCatalogTableOptimizerConfigurationOrphanFileDeletionConfigurationIcebergConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/glue_catalog_table_optimizer#location GlueCatalogTableOptimizer#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/glue_catalog_table_optimizer#orphan_file_retention_period_in_days GlueCatalogTableOptimizer#orphan_file_retention_period_in_days}.
	OrphanFileRetentionPeriodInDays *float64 `field:"optional" json:"orphanFileRetentionPeriodInDays" yaml:"orphanFileRetentionPeriodInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/glue_catalog_table_optimizer#run_rate_in_hours GlueCatalogTableOptimizer#run_rate_in_hours}.
	RunRateInHours *float64 `field:"optional" json:"runRateInHours" yaml:"runRateInHours"`
}

