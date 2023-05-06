package transferserver


type TransferServerWorkflowDetailsOnPartialUpload struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/transfer_server#execution_role TransferServer#execution_role}.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/transfer_server#workflow_id TransferServer#workflow_id}.
	WorkflowId *string `field:"required" json:"workflowId" yaml:"workflowId"`
}

