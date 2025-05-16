// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable


type GlueCatalogTableOpenTableFormatInputIcebergInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/glue_catalog_table#metadata_operation GlueCatalogTable#metadata_operation}.
	MetadataOperation *string `field:"required" json:"metadataOperation" yaml:"metadataOperation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/glue_catalog_table#version GlueCatalogTable#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

