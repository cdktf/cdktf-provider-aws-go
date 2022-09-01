package kinesis


type KinesisAnalyticsApplicationInputsProcessingConfiguration struct {
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_analytics_application#lambda KinesisAnalyticsApplication#lambda}
	Lambda *KinesisAnalyticsApplicationInputsProcessingConfigurationLambda `field:"required" json:"lambda" yaml:"lambda"`
}

