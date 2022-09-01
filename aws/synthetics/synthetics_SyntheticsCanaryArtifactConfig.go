package synthetics


type SyntheticsCanaryArtifactConfig struct {
	// s3_encryption block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/synthetics_canary#s3_encryption SyntheticsCanary#s3_encryption}
	S3Encryption *SyntheticsCanaryArtifactConfigS3Encryption `field:"optional" json:"s3Encryption" yaml:"s3Encryption"`
}

