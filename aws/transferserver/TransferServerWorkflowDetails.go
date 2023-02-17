package transferserver


type TransferServerWorkflowDetails struct {
	// on_upload block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_server#on_upload TransferServer#on_upload}
	OnUpload *TransferServerWorkflowDetailsOnUpload `field:"optional" json:"onUpload" yaml:"onUpload"`
}

