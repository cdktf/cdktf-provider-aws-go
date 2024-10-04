// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferserver


type TransferServerWorkflowDetails struct {
	// on_partial_upload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/transfer_server#on_partial_upload TransferServer#on_partial_upload}
	OnPartialUpload *TransferServerWorkflowDetailsOnPartialUpload `field:"optional" json:"onPartialUpload" yaml:"onPartialUpload"`
	// on_upload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/transfer_server#on_upload TransferServer#on_upload}
	OnUpload *TransferServerWorkflowDetailsOnUpload `field:"optional" json:"onUpload" yaml:"onUpload"`
}

