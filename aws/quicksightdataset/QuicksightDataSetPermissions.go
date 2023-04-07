package quicksightdataset


type QuicksightDataSetPermissions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_set#actions QuicksightDataSet#actions}.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_set#principal QuicksightDataSet#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

