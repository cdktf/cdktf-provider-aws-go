package s3bucket


type S3BucketReplicationConfigurationRulesDestinationAccessControlTranslation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket#owner S3Bucket#owner}.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
}

