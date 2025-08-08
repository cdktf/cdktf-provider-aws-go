// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable


type GlueCatalogTableStorageDescriptorSortColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/glue_catalog_table#column GlueCatalogTable#column}.
	Column *string `field:"required" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/glue_catalog_table#sort_order GlueCatalogTable#sort_order}.
	SortOrder *float64 `field:"required" json:"sortOrder" yaml:"sortOrder"`
}

