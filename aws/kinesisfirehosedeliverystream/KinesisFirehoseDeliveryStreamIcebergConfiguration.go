// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamIcebergConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#catalog_arn KinesisFirehoseDeliveryStream#catalog_arn}.
	CatalogArn *string `field:"required" json:"catalogArn" yaml:"catalogArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#role_arn KinesisFirehoseDeliveryStream#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#s3_configuration KinesisFirehoseDeliveryStream#s3_configuration}
	S3Configuration *KinesisFirehoseDeliveryStreamIcebergConfigurationS3Configuration `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#buffering_interval KinesisFirehoseDeliveryStream#buffering_interval}.
	BufferingInterval *float64 `field:"optional" json:"bufferingInterval" yaml:"bufferingInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#buffering_size KinesisFirehoseDeliveryStream#buffering_size}.
	BufferingSize *float64 `field:"optional" json:"bufferingSize" yaml:"bufferingSize"`
	// cloudwatch_logging_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#cloudwatch_logging_options KinesisFirehoseDeliveryStream#cloudwatch_logging_options}
	CloudwatchLoggingOptions *KinesisFirehoseDeliveryStreamIcebergConfigurationCloudwatchLoggingOptions `field:"optional" json:"cloudwatchLoggingOptions" yaml:"cloudwatchLoggingOptions"`
	// destination_table_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#destination_table_configuration KinesisFirehoseDeliveryStream#destination_table_configuration}
	DestinationTableConfiguration interface{} `field:"optional" json:"destinationTableConfiguration" yaml:"destinationTableConfiguration"`
	// processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#processing_configuration KinesisFirehoseDeliveryStream#processing_configuration}
	ProcessingConfiguration *KinesisFirehoseDeliveryStreamIcebergConfigurationProcessingConfiguration `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#retry_duration KinesisFirehoseDeliveryStream#retry_duration}.
	RetryDuration *float64 `field:"optional" json:"retryDuration" yaml:"retryDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_firehose_delivery_stream#s3_backup_mode KinesisFirehoseDeliveryStream#s3_backup_mode}.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
}

