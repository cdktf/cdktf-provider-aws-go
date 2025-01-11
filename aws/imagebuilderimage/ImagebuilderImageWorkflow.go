// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimage


type ImagebuilderImageWorkflow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/imagebuilder_image#workflow_arn ImagebuilderImage#workflow_arn}.
	WorkflowArn *string `field:"required" json:"workflowArn" yaml:"workflowArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/imagebuilder_image#on_failure ImagebuilderImage#on_failure}.
	OnFailure *string `field:"optional" json:"onFailure" yaml:"onFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/imagebuilder_image#parallel_group ImagebuilderImage#parallel_group}.
	ParallelGroup *string `field:"optional" json:"parallelGroup" yaml:"parallelGroup"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/imagebuilder_image#parameter ImagebuilderImage#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
}

