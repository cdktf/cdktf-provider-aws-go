// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtableoptimizer


type GlueCatalogTableOptimizerConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/glue_catalog_table_optimizer#enabled GlueCatalogTableOptimizer#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/glue_catalog_table_optimizer#role_arn GlueCatalogTableOptimizer#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// orphan_file_deletion_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/glue_catalog_table_optimizer#orphan_file_deletion_configuration GlueCatalogTableOptimizer#orphan_file_deletion_configuration}
	OrphanFileDeletionConfiguration interface{} `field:"optional" json:"orphanFileDeletionConfiguration" yaml:"orphanFileDeletionConfiguration"`
	// retention_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/glue_catalog_table_optimizer#retention_configuration GlueCatalogTableOptimizer#retention_configuration}
	RetentionConfiguration interface{} `field:"optional" json:"retentionConfiguration" yaml:"retentionConfiguration"`
}

