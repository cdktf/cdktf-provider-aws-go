// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferworkflow


type TransferWorkflowStepsCopyStepDetailsDestinationFileLocation struct {
	// efs_file_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/transfer_workflow#efs_file_location TransferWorkflow#efs_file_location}
	EfsFileLocation *TransferWorkflowStepsCopyStepDetailsDestinationFileLocationEfsFileLocation `field:"optional" json:"efsFileLocation" yaml:"efsFileLocation"`
	// s3_file_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/transfer_workflow#s3_file_location TransferWorkflow#s3_file_location}
	S3FileLocation *TransferWorkflowStepsCopyStepDetailsDestinationFileLocationS3FileLocation `field:"optional" json:"s3FileLocation" yaml:"s3FileLocation"`
}

