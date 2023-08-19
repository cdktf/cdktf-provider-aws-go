package syntheticscanary


type SyntheticsCanaryArtifactConfig struct {
	// s3_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/synthetics_canary#s3_encryption SyntheticsCanary#s3_encryption}
	S3Encryption *SyntheticsCanaryArtifactConfigS3Encryption `field:"optional" json:"s3Encryption" yaml:"s3Encryption"`
}

