package transfer


type TransferWorkflowOnExceptionStepsDeleteStepDetails struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_workflow#name TransferWorkflow#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_workflow#source_file_location TransferWorkflow#source_file_location}.
	SourceFileLocation *string `field:"optional" json:"sourceFileLocation" yaml:"sourceFileLocation"`
}

