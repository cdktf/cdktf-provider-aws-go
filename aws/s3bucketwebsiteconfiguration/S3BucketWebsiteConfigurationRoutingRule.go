package s3bucketwebsiteconfiguration


type S3BucketWebsiteConfigurationRoutingRule struct {
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/s3_bucket_website_configuration#redirect S3BucketWebsiteConfiguration#redirect}
	Redirect *S3BucketWebsiteConfigurationRoutingRuleRedirect `field:"required" json:"redirect" yaml:"redirect"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/s3_bucket_website_configuration#condition S3BucketWebsiteConfiguration#condition}
	Condition *S3BucketWebsiteConfigurationRoutingRuleCondition `field:"optional" json:"condition" yaml:"condition"`
}

