// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfiguration


type SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/sagemaker_endpoint_configuration#s3_output_path SagemakerEndpointConfiguration#s3_output_path}.
	S3OutputPath *string `field:"required" json:"s3OutputPath" yaml:"s3OutputPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/sagemaker_endpoint_configuration#kms_key_id SagemakerEndpointConfiguration#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// notification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/sagemaker_endpoint_configuration#notification_config SagemakerEndpointConfiguration#notification_config}
	NotificationConfig *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfig `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/sagemaker_endpoint_configuration#s3_failure_path SagemakerEndpointConfiguration#s3_failure_path}.
	S3FailurePath *string `field:"optional" json:"s3FailurePath" yaml:"s3FailurePath"`
}

