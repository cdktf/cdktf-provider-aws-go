// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimagepipeline


type ImagebuilderImagePipelineWorkflowParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/imagebuilder_image_pipeline#name ImagebuilderImagePipeline#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/imagebuilder_image_pipeline#value ImagebuilderImagePipeline#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

