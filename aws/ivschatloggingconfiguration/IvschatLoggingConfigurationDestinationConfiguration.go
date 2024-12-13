// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ivschatloggingconfiguration


type IvschatLoggingConfigurationDestinationConfiguration struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/ivschat_logging_configuration#cloudwatch_logs IvschatLoggingConfiguration#cloudwatch_logs}
	CloudwatchLogs *IvschatLoggingConfigurationDestinationConfigurationCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/ivschat_logging_configuration#firehose IvschatLoggingConfiguration#firehose}
	Firehose *IvschatLoggingConfigurationDestinationConfigurationFirehose `field:"optional" json:"firehose" yaml:"firehose"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/ivschat_logging_configuration#s3 IvschatLoggingConfiguration#s3}
	S3 *IvschatLoggingConfigurationDestinationConfigurationS3 `field:"optional" json:"s3" yaml:"s3"`
}

