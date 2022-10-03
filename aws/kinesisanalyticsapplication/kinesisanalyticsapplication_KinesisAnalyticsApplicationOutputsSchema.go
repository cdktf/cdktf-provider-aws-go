package kinesisanalyticsapplication


type KinesisAnalyticsApplicationOutputsSchema struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kinesis_analytics_application#record_format_type KinesisAnalyticsApplication#record_format_type}.
	RecordFormatType *string `field:"required" json:"recordFormatType" yaml:"recordFormatType"`
}

