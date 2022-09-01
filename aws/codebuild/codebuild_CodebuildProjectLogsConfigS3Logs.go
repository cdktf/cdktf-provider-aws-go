package codebuild


type CodebuildProjectLogsConfigS3Logs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#bucket_owner_access CodebuildProject#bucket_owner_access}.
	BucketOwnerAccess *string `field:"optional" json:"bucketOwnerAccess" yaml:"bucketOwnerAccess"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#encryption_disabled CodebuildProject#encryption_disabled}.
	EncryptionDisabled interface{} `field:"optional" json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#location CodebuildProject#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#status CodebuildProject#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

