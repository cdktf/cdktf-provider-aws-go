package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsTagColumnOperationTags struct {
	// column_description block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/quicksight_data_set#column_description QuicksightDataSet#column_description}
	ColumnDescription *QuicksightDataSetLogicalTableMapDataTransformsTagColumnOperationTagsColumnDescription `field:"optional" json:"columnDescription" yaml:"columnDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/quicksight_data_set#column_geographic_role QuicksightDataSet#column_geographic_role}.
	ColumnGeographicRole *string `field:"optional" json:"columnGeographicRole" yaml:"columnGeographicRole"`
}

