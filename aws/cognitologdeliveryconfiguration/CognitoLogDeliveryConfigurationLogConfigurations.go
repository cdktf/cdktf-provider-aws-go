// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitologdeliveryconfiguration


type CognitoLogDeliveryConfigurationLogConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/cognito_log_delivery_configuration#event_source CognitoLogDeliveryConfiguration#event_source}.
	EventSource *string `field:"required" json:"eventSource" yaml:"eventSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/cognito_log_delivery_configuration#log_level CognitoLogDeliveryConfiguration#log_level}.
	LogLevel *string `field:"required" json:"logLevel" yaml:"logLevel"`
	// cloud_watch_logs_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/cognito_log_delivery_configuration#cloud_watch_logs_configuration CognitoLogDeliveryConfiguration#cloud_watch_logs_configuration}
	CloudWatchLogsConfiguration interface{} `field:"optional" json:"cloudWatchLogsConfiguration" yaml:"cloudWatchLogsConfiguration"`
	// firehose_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/cognito_log_delivery_configuration#firehose_configuration CognitoLogDeliveryConfiguration#firehose_configuration}
	FirehoseConfiguration interface{} `field:"optional" json:"firehoseConfiguration" yaml:"firehoseConfiguration"`
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/cognito_log_delivery_configuration#s3_configuration CognitoLogDeliveryConfiguration#s3_configuration}
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

