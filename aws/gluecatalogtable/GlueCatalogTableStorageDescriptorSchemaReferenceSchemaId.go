// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable


type GlueCatalogTableStorageDescriptorSchemaReferenceSchemaId struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/glue_catalog_table#registry_name GlueCatalogTable#registry_name}.
	RegistryName *string `field:"optional" json:"registryName" yaml:"registryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/glue_catalog_table#schema_arn GlueCatalogTable#schema_arn}.
	SchemaArn *string `field:"optional" json:"schemaArn" yaml:"schemaArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/glue_catalog_table#schema_name GlueCatalogTable#schema_name}.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
}

