// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtableoptimizer


type GlueCatalogTableOptimizerConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/glue_catalog_table_optimizer#enabled GlueCatalogTableOptimizer#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/glue_catalog_table_optimizer#role_arn GlueCatalogTableOptimizer#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

