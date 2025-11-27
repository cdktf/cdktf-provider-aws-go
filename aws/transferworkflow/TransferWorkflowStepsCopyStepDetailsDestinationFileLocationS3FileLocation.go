// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferworkflow


type TransferWorkflowStepsCopyStepDetailsDestinationFileLocationS3FileLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/transfer_workflow#bucket TransferWorkflow#bucket}.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/transfer_workflow#key TransferWorkflow#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

