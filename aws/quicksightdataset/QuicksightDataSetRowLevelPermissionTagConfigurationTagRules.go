package quicksightdataset


type QuicksightDataSetRowLevelPermissionTagConfigurationTagRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/quicksight_data_set#column_name QuicksightDataSet#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/quicksight_data_set#tag_key QuicksightDataSet#tag_key}.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/quicksight_data_set#match_all_value QuicksightDataSet#match_all_value}.
	MatchAllValue *string `field:"optional" json:"matchAllValue" yaml:"matchAllValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/quicksight_data_set#tag_multi_value_delimiter QuicksightDataSet#tag_multi_value_delimiter}.
	TagMultiValueDelimiter *string `field:"optional" json:"tagMultiValueDelimiter" yaml:"tagMultiValueDelimiter"`
}

