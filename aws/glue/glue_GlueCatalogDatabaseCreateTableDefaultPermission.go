package glue


type GlueCatalogDatabaseCreateTableDefaultPermission struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_database#permissions GlueCatalogDatabase#permissions}.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// principal block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_database#principal GlueCatalogDatabase#principal}
	Principal *GlueCatalogDatabaseCreateTableDefaultPermissionPrincipal `field:"optional" json:"principal" yaml:"principal"`
}

