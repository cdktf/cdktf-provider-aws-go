package s3bucketversioning


type S3BucketVersioningVersioningConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/s3_bucket_versioning#status S3BucketVersioningA#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/s3_bucket_versioning#mfa_delete S3BucketVersioningA#mfa_delete}.
	MfaDelete *string `field:"optional" json:"mfaDelete" yaml:"mfaDelete"`
}

