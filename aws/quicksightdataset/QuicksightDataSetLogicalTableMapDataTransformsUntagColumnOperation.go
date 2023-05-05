package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsUntagColumnOperation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/quicksight_data_set#column_name QuicksightDataSet#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/quicksight_data_set#tag_names QuicksightDataSet#tag_names}.
	TagNames *[]*string `field:"required" json:"tagNames" yaml:"tagNames"`
}

