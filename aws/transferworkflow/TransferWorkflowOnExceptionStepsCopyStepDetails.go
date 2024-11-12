// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferworkflow


type TransferWorkflowOnExceptionStepsCopyStepDetails struct {
	// destination_file_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/transfer_workflow#destination_file_location TransferWorkflow#destination_file_location}
	DestinationFileLocation *TransferWorkflowOnExceptionStepsCopyStepDetailsDestinationFileLocation `field:"optional" json:"destinationFileLocation" yaml:"destinationFileLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/transfer_workflow#name TransferWorkflow#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/transfer_workflow#overwrite_existing TransferWorkflow#overwrite_existing}.
	OverwriteExisting *string `field:"optional" json:"overwriteExisting" yaml:"overwriteExisting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/transfer_workflow#source_file_location TransferWorkflow#source_file_location}.
	SourceFileLocation *string `field:"optional" json:"sourceFileLocation" yaml:"sourceFileLocation"`
}

