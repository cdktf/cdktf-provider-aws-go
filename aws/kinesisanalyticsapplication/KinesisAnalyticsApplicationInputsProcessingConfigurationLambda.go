package kinesisanalyticsapplication


type KinesisAnalyticsApplicationInputsProcessingConfigurationLambda struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/kinesis_analytics_application#resource_arn KinesisAnalyticsApplication#resource_arn}.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/kinesis_analytics_application#role_arn KinesisAnalyticsApplication#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

