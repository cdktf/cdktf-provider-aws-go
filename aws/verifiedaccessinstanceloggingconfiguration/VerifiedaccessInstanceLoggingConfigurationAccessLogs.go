// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedaccessinstanceloggingconfiguration


type VerifiedaccessInstanceLoggingConfigurationAccessLogs struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/verifiedaccess_instance_logging_configuration#cloudwatch_logs VerifiedaccessInstanceLoggingConfiguration#cloudwatch_logs}
	CloudwatchLogs *VerifiedaccessInstanceLoggingConfigurationAccessLogsCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/verifiedaccess_instance_logging_configuration#include_trust_context VerifiedaccessInstanceLoggingConfiguration#include_trust_context}.
	IncludeTrustContext interface{} `field:"optional" json:"includeTrustContext" yaml:"includeTrustContext"`
	// kinesis_data_firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/verifiedaccess_instance_logging_configuration#kinesis_data_firehose VerifiedaccessInstanceLoggingConfiguration#kinesis_data_firehose}
	KinesisDataFirehose *VerifiedaccessInstanceLoggingConfigurationAccessLogsKinesisDataFirehose `field:"optional" json:"kinesisDataFirehose" yaml:"kinesisDataFirehose"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/verifiedaccess_instance_logging_configuration#log_version VerifiedaccessInstanceLoggingConfiguration#log_version}.
	LogVersion *string `field:"optional" json:"logVersion" yaml:"logVersion"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/verifiedaccess_instance_logging_configuration#s3 VerifiedaccessInstanceLoggingConfiguration#s3}
	S3 *VerifiedaccessInstanceLoggingConfigurationAccessLogsS3 `field:"optional" json:"s3" yaml:"s3"`
}

