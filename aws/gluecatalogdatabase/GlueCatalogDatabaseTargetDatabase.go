package gluecatalogdatabase


type GlueCatalogDatabaseTargetDatabase struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/glue_catalog_database#catalog_id GlueCatalogDatabase#catalog_id}.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/glue_catalog_database#database_name GlueCatalogDatabase#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
}

