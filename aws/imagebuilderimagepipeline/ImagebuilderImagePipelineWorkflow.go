// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimagepipeline


type ImagebuilderImagePipelineWorkflow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/imagebuilder_image_pipeline#workflow_arn ImagebuilderImagePipeline#workflow_arn}.
	WorkflowArn *string `field:"required" json:"workflowArn" yaml:"workflowArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/imagebuilder_image_pipeline#on_failure ImagebuilderImagePipeline#on_failure}.
	OnFailure *string `field:"optional" json:"onFailure" yaml:"onFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/imagebuilder_image_pipeline#parallel_group ImagebuilderImagePipeline#parallel_group}.
	ParallelGroup *string `field:"optional" json:"parallelGroup" yaml:"parallelGroup"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/imagebuilder_image_pipeline#parameter ImagebuilderImagePipeline#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
}

