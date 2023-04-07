package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsRenameColumnOperation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_set#column_name QuicksightDataSet#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_set#new_column_name QuicksightDataSet#new_column_name}.
	NewColumnName *string `field:"required" json:"newColumnName" yaml:"newColumnName"`
}

