package quicksightdataset


type QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/quicksight_data_set#column_name QuicksightDataSet#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/quicksight_data_set#size QuicksightDataSet#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/quicksight_data_set#size_unit QuicksightDataSet#size_unit}.
	SizeUnit *string `field:"required" json:"sizeUnit" yaml:"sizeUnit"`
}

