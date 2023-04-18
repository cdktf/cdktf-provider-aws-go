package kinesisanalyticsapplication


type KinesisAnalyticsApplicationReferenceDataSources struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/kinesis_analytics_application#s3 KinesisAnalyticsApplication#s3}
	S3 *KinesisAnalyticsApplicationReferenceDataSourcesS3 `field:"required" json:"s3" yaml:"s3"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/kinesis_analytics_application#schema KinesisAnalyticsApplication#schema}
	Schema *KinesisAnalyticsApplicationReferenceDataSourcesSchema `field:"required" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/kinesis_analytics_application#table_name KinesisAnalyticsApplication#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

