// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtableoptimizer


type GlueCatalogTableOptimizerConfigurationRetentionConfiguration struct {
	// iceberg_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/glue_catalog_table_optimizer#iceberg_configuration GlueCatalogTableOptimizer#iceberg_configuration}
	IcebergConfiguration interface{} `field:"optional" json:"icebergConfiguration" yaml:"icebergConfiguration"`
}

