package s3bucketacl


type S3BucketAclAccessControlPolicyGrant struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/s3_bucket_acl#permission S3BucketAcl#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// grantee block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/s3_bucket_acl#grantee S3BucketAcl#grantee}
	Grantee *S3BucketAclAccessControlPolicyGrantGrantee `field:"optional" json:"grantee" yaml:"grantee"`
}

