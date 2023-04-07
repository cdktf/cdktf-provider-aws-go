package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsUntagColumnOperation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_set#column_name QuicksightDataSet#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_set#tag_names QuicksightDataSet#tag_names}.
	TagNames *[]*string `field:"required" json:"tagNames" yaml:"tagNames"`
}

