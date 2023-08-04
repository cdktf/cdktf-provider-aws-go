package quicksightdataset


type QuicksightDataSetColumnLevelPermissionRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/quicksight_data_set#column_names QuicksightDataSet#column_names}.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/quicksight_data_set#principals QuicksightDataSet#principals}.
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
}

