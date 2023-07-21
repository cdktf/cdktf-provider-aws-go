package quicksightdataset


type QuicksightDataSetPhysicalTableMapCustomSql struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/quicksight_data_set#data_source_arn QuicksightDataSet#data_source_arn}.
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/quicksight_data_set#name QuicksightDataSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/quicksight_data_set#sql_query QuicksightDataSet#sql_query}.
	SqlQuery *string `field:"required" json:"sqlQuery" yaml:"sqlQuery"`
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/quicksight_data_set#columns QuicksightDataSet#columns}
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
}

