package kinesis


type KinesisAnalyticsApplicationInputsSchemaRecordFormatMappingParameters struct {
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_analytics_application#csv KinesisAnalyticsApplication#csv}
	Csv *KinesisAnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsv `field:"optional" json:"csv" yaml:"csv"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_analytics_application#json KinesisAnalyticsApplication#json}
	Json *KinesisAnalyticsApplicationInputsSchemaRecordFormatMappingParametersJson `field:"optional" json:"json" yaml:"json"`
}

