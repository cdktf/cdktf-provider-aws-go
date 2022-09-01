package kinesis


type Kinesisanalyticsv2ApplicationCloudwatchLoggingOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application#log_stream_arn Kinesisanalyticsv2Application#log_stream_arn}.
	LogStreamArn *string `field:"required" json:"logStreamArn" yaml:"logStreamArn"`
}

