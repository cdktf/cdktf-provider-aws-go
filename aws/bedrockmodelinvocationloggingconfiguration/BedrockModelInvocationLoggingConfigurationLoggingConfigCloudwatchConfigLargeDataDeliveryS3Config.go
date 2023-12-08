// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockmodelinvocationloggingconfiguration


type BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3Config struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/bedrock_model_invocation_logging_configuration#bucket_name BedrockModelInvocationLoggingConfiguration#bucket_name}.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/bedrock_model_invocation_logging_configuration#key_prefix BedrockModelInvocationLoggingConfiguration#key_prefix}.
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

