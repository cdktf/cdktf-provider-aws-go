package s3controlbucketlifecycleconfiguration


type S3ControlBucketLifecycleConfigurationRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/s3control_bucket_lifecycle_configuration#id S3ControlBucketLifecycleConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// abort_incomplete_multipart_upload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/s3control_bucket_lifecycle_configuration#abort_incomplete_multipart_upload S3ControlBucketLifecycleConfiguration#abort_incomplete_multipart_upload}
	AbortIncompleteMultipartUpload *S3ControlBucketLifecycleConfigurationRuleAbortIncompleteMultipartUpload `field:"optional" json:"abortIncompleteMultipartUpload" yaml:"abortIncompleteMultipartUpload"`
	// expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/s3control_bucket_lifecycle_configuration#expiration S3ControlBucketLifecycleConfiguration#expiration}
	Expiration *S3ControlBucketLifecycleConfigurationRuleExpiration `field:"optional" json:"expiration" yaml:"expiration"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/s3control_bucket_lifecycle_configuration#filter S3ControlBucketLifecycleConfiguration#filter}
	Filter *S3ControlBucketLifecycleConfigurationRuleFilter `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/s3control_bucket_lifecycle_configuration#status S3ControlBucketLifecycleConfiguration#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

