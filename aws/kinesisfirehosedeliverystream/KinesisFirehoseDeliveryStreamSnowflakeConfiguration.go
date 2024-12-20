// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamSnowflakeConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#account_url KinesisFirehoseDeliveryStream#account_url}.
	AccountUrl *string `field:"required" json:"accountUrl" yaml:"accountUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#database KinesisFirehoseDeliveryStream#database}.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#role_arn KinesisFirehoseDeliveryStream#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#s3_configuration KinesisFirehoseDeliveryStream#s3_configuration}
	S3Configuration *KinesisFirehoseDeliveryStreamSnowflakeConfigurationS3Configuration `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#schema KinesisFirehoseDeliveryStream#schema}.
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#table KinesisFirehoseDeliveryStream#table}.
	Table *string `field:"required" json:"table" yaml:"table"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#buffering_interval KinesisFirehoseDeliveryStream#buffering_interval}.
	BufferingInterval *float64 `field:"optional" json:"bufferingInterval" yaml:"bufferingInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#buffering_size KinesisFirehoseDeliveryStream#buffering_size}.
	BufferingSize *float64 `field:"optional" json:"bufferingSize" yaml:"bufferingSize"`
	// cloudwatch_logging_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#cloudwatch_logging_options KinesisFirehoseDeliveryStream#cloudwatch_logging_options}
	CloudwatchLoggingOptions *KinesisFirehoseDeliveryStreamSnowflakeConfigurationCloudwatchLoggingOptions `field:"optional" json:"cloudwatchLoggingOptions" yaml:"cloudwatchLoggingOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#content_column_name KinesisFirehoseDeliveryStream#content_column_name}.
	ContentColumnName *string `field:"optional" json:"contentColumnName" yaml:"contentColumnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#data_loading_option KinesisFirehoseDeliveryStream#data_loading_option}.
	DataLoadingOption *string `field:"optional" json:"dataLoadingOption" yaml:"dataLoadingOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#key_passphrase KinesisFirehoseDeliveryStream#key_passphrase}.
	KeyPassphrase *string `field:"optional" json:"keyPassphrase" yaml:"keyPassphrase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#metadata_column_name KinesisFirehoseDeliveryStream#metadata_column_name}.
	MetadataColumnName *string `field:"optional" json:"metadataColumnName" yaml:"metadataColumnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#private_key KinesisFirehoseDeliveryStream#private_key}.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#processing_configuration KinesisFirehoseDeliveryStream#processing_configuration}
	ProcessingConfiguration *KinesisFirehoseDeliveryStreamSnowflakeConfigurationProcessingConfiguration `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#retry_duration KinesisFirehoseDeliveryStream#retry_duration}.
	RetryDuration *float64 `field:"optional" json:"retryDuration" yaml:"retryDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#s3_backup_mode KinesisFirehoseDeliveryStream#s3_backup_mode}.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
	// secrets_manager_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#secrets_manager_configuration KinesisFirehoseDeliveryStream#secrets_manager_configuration}
	SecretsManagerConfiguration *KinesisFirehoseDeliveryStreamSnowflakeConfigurationSecretsManagerConfiguration `field:"optional" json:"secretsManagerConfiguration" yaml:"secretsManagerConfiguration"`
	// snowflake_role_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#snowflake_role_configuration KinesisFirehoseDeliveryStream#snowflake_role_configuration}
	SnowflakeRoleConfiguration *KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeRoleConfiguration `field:"optional" json:"snowflakeRoleConfiguration" yaml:"snowflakeRoleConfiguration"`
	// snowflake_vpc_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#snowflake_vpc_configuration KinesisFirehoseDeliveryStream#snowflake_vpc_configuration}
	SnowflakeVpcConfiguration *KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeVpcConfiguration `field:"optional" json:"snowflakeVpcConfiguration" yaml:"snowflakeVpcConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/kinesis_firehose_delivery_stream#user KinesisFirehoseDeliveryStream#user}.
	User *string `field:"optional" json:"user" yaml:"user"`
}

