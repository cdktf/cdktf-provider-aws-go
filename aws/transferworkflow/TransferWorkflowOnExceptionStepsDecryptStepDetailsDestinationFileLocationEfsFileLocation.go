// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferworkflow


type TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/transfer_workflow#file_system_id TransferWorkflow#file_system_id}.
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/transfer_workflow#path TransferWorkflow#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

