// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfiguration


type SagemakerEndpointConfigurationDataCaptureConfig struct {
	// capture_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/sagemaker_endpoint_configuration#capture_options SagemakerEndpointConfiguration#capture_options}
	CaptureOptions interface{} `field:"required" json:"captureOptions" yaml:"captureOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/sagemaker_endpoint_configuration#destination_s3_uri SagemakerEndpointConfiguration#destination_s3_uri}.
	DestinationS3Uri *string `field:"required" json:"destinationS3Uri" yaml:"destinationS3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/sagemaker_endpoint_configuration#initial_sampling_percentage SagemakerEndpointConfiguration#initial_sampling_percentage}.
	InitialSamplingPercentage *float64 `field:"required" json:"initialSamplingPercentage" yaml:"initialSamplingPercentage"`
	// capture_content_type_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/sagemaker_endpoint_configuration#capture_content_type_header SagemakerEndpointConfiguration#capture_content_type_header}
	CaptureContentTypeHeader *SagemakerEndpointConfigurationDataCaptureConfigCaptureContentTypeHeader `field:"optional" json:"captureContentTypeHeader" yaml:"captureContentTypeHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/sagemaker_endpoint_configuration#enable_capture SagemakerEndpointConfiguration#enable_capture}.
	EnableCapture interface{} `field:"optional" json:"enableCapture" yaml:"enableCapture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/sagemaker_endpoint_configuration#kms_key_id SagemakerEndpointConfiguration#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

