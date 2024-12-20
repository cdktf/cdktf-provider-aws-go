// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable


type GlueCatalogTableStorageDescriptorSerDeInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/glue_catalog_table#name GlueCatalogTable#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/glue_catalog_table#parameters GlueCatalogTable#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/glue_catalog_table#serialization_library GlueCatalogTable#serialization_library}.
	SerializationLibrary *string `field:"optional" json:"serializationLibrary" yaml:"serializationLibrary"`
}

