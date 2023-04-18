package transferserver


type TransferServerWorkflowDetails struct {
	// on_partial_upload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/transfer_server#on_partial_upload TransferServer#on_partial_upload}
	OnPartialUpload *TransferServerWorkflowDetailsOnPartialUpload `field:"optional" json:"onPartialUpload" yaml:"onPartialUpload"`
	// on_upload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/transfer_server#on_upload TransferServer#on_upload}
	OnUpload *TransferServerWorkflowDetailsOnUpload `field:"optional" json:"onUpload" yaml:"onUpload"`
}

