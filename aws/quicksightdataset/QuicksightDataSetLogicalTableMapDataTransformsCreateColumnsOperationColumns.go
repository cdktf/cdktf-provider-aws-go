package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsCreateColumnsOperationColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/quicksight_data_set#column_id QuicksightDataSet#column_id}.
	ColumnId *string `field:"required" json:"columnId" yaml:"columnId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/quicksight_data_set#column_name QuicksightDataSet#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/quicksight_data_set#expression QuicksightDataSet#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

