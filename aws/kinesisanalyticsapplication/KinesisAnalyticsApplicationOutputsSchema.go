package kinesisanalyticsapplication


type KinesisAnalyticsApplicationOutputsSchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/kinesis_analytics_application#record_format_type KinesisAnalyticsApplication#record_format_type}.
	RecordFormatType *string `field:"required" json:"recordFormatType" yaml:"recordFormatType"`
}

