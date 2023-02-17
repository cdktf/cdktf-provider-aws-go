package transferworkflow


type TransferWorkflowOnExceptionStepsCopyStepDetailsDestinationFileLocationS3FileLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_workflow#bucket TransferWorkflow#bucket}.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_workflow#key TransferWorkflow#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

