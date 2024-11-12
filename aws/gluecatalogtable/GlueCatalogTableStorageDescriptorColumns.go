// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable


type GlueCatalogTableStorageDescriptorColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/glue_catalog_table#name GlueCatalogTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/glue_catalog_table#comment GlueCatalogTable#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/glue_catalog_table#parameters GlueCatalogTable#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/glue_catalog_table#type GlueCatalogTable#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

