package kinesis


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfiguration struct {
	// input block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application#input Kinesisanalyticsv2Application#input}
	Input *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInput `field:"optional" json:"input" yaml:"input"`
	// output block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application#output Kinesisanalyticsv2Application#output}
	Output interface{} `field:"optional" json:"output" yaml:"output"`
	// reference_data_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application#reference_data_source Kinesisanalyticsv2Application#reference_data_source}
	ReferenceDataSource *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSource `field:"optional" json:"referenceDataSource" yaml:"referenceDataSource"`
}

