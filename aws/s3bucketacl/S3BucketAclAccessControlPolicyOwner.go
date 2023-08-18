package s3bucketacl


type S3BucketAclAccessControlPolicyOwner struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/s3_bucket_acl#id S3BucketAcl#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/s3_bucket_acl#display_name S3BucketAcl#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

