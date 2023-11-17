// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamSplunkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/kinesis_firehose_delivery_stream#hec_endpoint KinesisFirehoseDeliveryStream#hec_endpoint}.
	HecEndpoint *string `field:"required" json:"hecEndpoint" yaml:"hecEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/kinesis_firehose_delivery_stream#hec_token KinesisFirehoseDeliveryStream#hec_token}.
	HecToken *string `field:"required" json:"hecToken" yaml:"hecToken"`
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/kinesis_firehose_delivery_stream#s3_configuration KinesisFirehoseDeliveryStream#s3_configuration}
	S3Configuration *KinesisFirehoseDeliveryStreamSplunkConfigurationS3Configuration `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// cloudwatch_logging_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/kinesis_firehose_delivery_stream#cloudwatch_logging_options KinesisFirehoseDeliveryStream#cloudwatch_logging_options}
	CloudwatchLoggingOptions *KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptions `field:"optional" json:"cloudwatchLoggingOptions" yaml:"cloudwatchLoggingOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/kinesis_firehose_delivery_stream#hec_acknowledgment_timeout KinesisFirehoseDeliveryStream#hec_acknowledgment_timeout}.
	HecAcknowledgmentTimeout *float64 `field:"optional" json:"hecAcknowledgmentTimeout" yaml:"hecAcknowledgmentTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/kinesis_firehose_delivery_stream#hec_endpoint_type KinesisFirehoseDeliveryStream#hec_endpoint_type}.
	HecEndpointType *string `field:"optional" json:"hecEndpointType" yaml:"hecEndpointType"`
	// processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/kinesis_firehose_delivery_stream#processing_configuration KinesisFirehoseDeliveryStream#processing_configuration}
	ProcessingConfiguration *KinesisFirehoseDeliveryStreamSplunkConfigurationProcessingConfiguration `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/kinesis_firehose_delivery_stream#retry_duration KinesisFirehoseDeliveryStream#retry_duration}.
	RetryDuration *float64 `field:"optional" json:"retryDuration" yaml:"retryDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/kinesis_firehose_delivery_stream#s3_backup_mode KinesisFirehoseDeliveryStream#s3_backup_mode}.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
}

