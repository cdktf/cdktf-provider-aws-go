package s3bucket


type S3BucketLifecycleRuleNoncurrentVersionExpiration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/s3_bucket#days S3Bucket#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

