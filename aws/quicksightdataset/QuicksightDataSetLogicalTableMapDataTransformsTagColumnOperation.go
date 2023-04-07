package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsTagColumnOperation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_set#column_name QuicksightDataSet#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_set#tags QuicksightDataSet#tags}
	Tags interface{} `field:"required" json:"tags" yaml:"tags"`
}

