// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferworkflow


type TransferWorkflowOnExceptionStepsTagStepDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/transfer_workflow#name TransferWorkflow#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/transfer_workflow#source_file_location TransferWorkflow#source_file_location}.
	SourceFileLocation *string `field:"optional" json:"sourceFileLocation" yaml:"sourceFileLocation"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/transfer_workflow#tags TransferWorkflow#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

