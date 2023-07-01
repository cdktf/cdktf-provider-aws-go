package kinesisanalyticsapplication


type KinesisAnalyticsApplicationInputsProcessingConfiguration struct {
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/kinesis_analytics_application#lambda KinesisAnalyticsApplication#lambda}
	Lambda *KinesisAnalyticsApplicationInputsProcessingConfigurationLambda `field:"required" json:"lambda" yaml:"lambda"`
}

