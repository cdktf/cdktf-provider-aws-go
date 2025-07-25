// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamRedshiftConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#cluster_jdbcurl KinesisFirehoseDeliveryStream#cluster_jdbcurl}.
	ClusterJdbcurl *string `field:"required" json:"clusterJdbcurl" yaml:"clusterJdbcurl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#data_table_name KinesisFirehoseDeliveryStream#data_table_name}.
	DataTableName *string `field:"required" json:"dataTableName" yaml:"dataTableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#role_arn KinesisFirehoseDeliveryStream#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#s3_configuration KinesisFirehoseDeliveryStream#s3_configuration}
	S3Configuration *KinesisFirehoseDeliveryStreamRedshiftConfigurationS3Configuration `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// cloudwatch_logging_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#cloudwatch_logging_options KinesisFirehoseDeliveryStream#cloudwatch_logging_options}
	CloudwatchLoggingOptions *KinesisFirehoseDeliveryStreamRedshiftConfigurationCloudwatchLoggingOptions `field:"optional" json:"cloudwatchLoggingOptions" yaml:"cloudwatchLoggingOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#copy_options KinesisFirehoseDeliveryStream#copy_options}.
	CopyOptions *string `field:"optional" json:"copyOptions" yaml:"copyOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#data_table_columns KinesisFirehoseDeliveryStream#data_table_columns}.
	DataTableColumns *string `field:"optional" json:"dataTableColumns" yaml:"dataTableColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#password KinesisFirehoseDeliveryStream#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#processing_configuration KinesisFirehoseDeliveryStream#processing_configuration}
	ProcessingConfiguration *KinesisFirehoseDeliveryStreamRedshiftConfigurationProcessingConfiguration `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#retry_duration KinesisFirehoseDeliveryStream#retry_duration}.
	RetryDuration *float64 `field:"optional" json:"retryDuration" yaml:"retryDuration"`
	// s3_backup_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#s3_backup_configuration KinesisFirehoseDeliveryStream#s3_backup_configuration}
	S3BackupConfiguration *KinesisFirehoseDeliveryStreamRedshiftConfigurationS3BackupConfiguration `field:"optional" json:"s3BackupConfiguration" yaml:"s3BackupConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#s3_backup_mode KinesisFirehoseDeliveryStream#s3_backup_mode}.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
	// secrets_manager_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#secrets_manager_configuration KinesisFirehoseDeliveryStream#secrets_manager_configuration}
	SecretsManagerConfiguration *KinesisFirehoseDeliveryStreamRedshiftConfigurationSecretsManagerConfiguration `field:"optional" json:"secretsManagerConfiguration" yaml:"secretsManagerConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#username KinesisFirehoseDeliveryStream#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

