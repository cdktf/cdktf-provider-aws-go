package gluecatalogtable


type GlueCatalogTableStorageDescriptorSortColumns struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table#column GlueCatalogTable#column}.
	Column *string `field:"required" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_catalog_table#sort_order GlueCatalogTable#sort_order}.
	SortOrder *float64 `field:"required" json:"sortOrder" yaml:"sortOrder"`
}

