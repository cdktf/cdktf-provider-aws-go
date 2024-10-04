// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockmodelinvocationloggingconfiguration


type BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfig struct {
	// large_data_delivery_s3_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/bedrock_model_invocation_logging_configuration#large_data_delivery_s3_config BedrockModelInvocationLoggingConfiguration#large_data_delivery_s3_config}
	LargeDataDeliveryS3Config *BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3Config `field:"optional" json:"largeDataDeliveryS3Config" yaml:"largeDataDeliveryS3Config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/bedrock_model_invocation_logging_configuration#log_group_name BedrockModelInvocationLoggingConfiguration#log_group_name}.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/bedrock_model_invocation_logging_configuration#role_arn BedrockModelInvocationLoggingConfiguration#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

