package kinesis


type KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormat struct {
	// mapping_parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_analytics_application#mapping_parameters KinesisAnalyticsApplication#mapping_parameters}
	MappingParameters *KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatMappingParameters `field:"optional" json:"mappingParameters" yaml:"mappingParameters"`
}

