package quicksightdataset


type QuicksightDataSetPhysicalTableMapS3SourceInputColumns struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_set#name QuicksightDataSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_set#type QuicksightDataSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

