package kinesis


type KinesisAnalyticsApplicationCloudwatchLoggingOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_analytics_application#log_stream_arn KinesisAnalyticsApplication#log_stream_arn}.
	LogStreamArn *string `field:"required" json:"logStreamArn" yaml:"logStreamArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_analytics_application#role_arn KinesisAnalyticsApplication#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

