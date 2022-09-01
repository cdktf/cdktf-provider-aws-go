package kinesis


type KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatMappingParameters struct {
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_analytics_application#csv KinesisAnalyticsApplication#csv}
	Csv *KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatMappingParametersCsv `field:"optional" json:"csv" yaml:"csv"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_analytics_application#json KinesisAnalyticsApplication#json}
	Json *KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatMappingParametersJson `field:"optional" json:"json" yaml:"json"`
}

