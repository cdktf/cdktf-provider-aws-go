package quicksightdataset


type QuicksightDataSetColumnLevelPermissionRules struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_set#column_names QuicksightDataSet#column_names}.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_set#principals QuicksightDataSet#principals}.
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
}

