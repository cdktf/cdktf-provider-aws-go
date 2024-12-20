// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerpipeline


type SagemakerPipelinePipelineDefinitionS3Location struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/sagemaker_pipeline#bucket SagemakerPipeline#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/sagemaker_pipeline#object_key SagemakerPipeline#object_key}.
	ObjectKey *string `field:"required" json:"objectKey" yaml:"objectKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/sagemaker_pipeline#version_id SagemakerPipeline#version_id}.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

