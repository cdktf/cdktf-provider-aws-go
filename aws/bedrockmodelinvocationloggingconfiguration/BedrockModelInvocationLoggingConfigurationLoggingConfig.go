// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockmodelinvocationloggingconfiguration


type BedrockModelInvocationLoggingConfigurationLoggingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/bedrock_model_invocation_logging_configuration#embedding_data_delivery_enabled BedrockModelInvocationLoggingConfiguration#embedding_data_delivery_enabled}.
	EmbeddingDataDeliveryEnabled interface{} `field:"required" json:"embeddingDataDeliveryEnabled" yaml:"embeddingDataDeliveryEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/bedrock_model_invocation_logging_configuration#image_data_delivery_enabled BedrockModelInvocationLoggingConfiguration#image_data_delivery_enabled}.
	ImageDataDeliveryEnabled interface{} `field:"required" json:"imageDataDeliveryEnabled" yaml:"imageDataDeliveryEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/bedrock_model_invocation_logging_configuration#text_data_delivery_enabled BedrockModelInvocationLoggingConfiguration#text_data_delivery_enabled}.
	TextDataDeliveryEnabled interface{} `field:"required" json:"textDataDeliveryEnabled" yaml:"textDataDeliveryEnabled"`
	// cloudwatch_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/bedrock_model_invocation_logging_configuration#cloudwatch_config BedrockModelInvocationLoggingConfiguration#cloudwatch_config}
	CloudwatchConfig *BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfig `field:"optional" json:"cloudwatchConfig" yaml:"cloudwatchConfig"`
	// s3_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/bedrock_model_invocation_logging_configuration#s3_config BedrockModelInvocationLoggingConfiguration#s3_config}
	S3Config *BedrockModelInvocationLoggingConfigurationLoggingConfigS3Config `field:"optional" json:"s3Config" yaml:"s3Config"`
}

