package s3


type S3BucketAclAccessControlPolicyGrant struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_acl#permission S3BucketAcl#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// grantee block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_acl#grantee S3BucketAcl#grantee}
	Grantee *S3BucketAclAccessControlPolicyGrantGrantee `field:"optional" json:"grantee" yaml:"grantee"`
}

