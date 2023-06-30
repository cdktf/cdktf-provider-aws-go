package gluecatalogtable


type GlueCatalogTableStorageDescriptorSchemaReference struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/glue_catalog_table#schema_version_number GlueCatalogTable#schema_version_number}.
	SchemaVersionNumber *float64 `field:"required" json:"schemaVersionNumber" yaml:"schemaVersionNumber"`
	// schema_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/glue_catalog_table#schema_id GlueCatalogTable#schema_id}
	SchemaId *GlueCatalogTableStorageDescriptorSchemaReferenceSchemaId `field:"optional" json:"schemaId" yaml:"schemaId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/glue_catalog_table#schema_version_id GlueCatalogTable#schema_version_id}.
	SchemaVersionId *string `field:"optional" json:"schemaVersionId" yaml:"schemaVersionId"`
}

