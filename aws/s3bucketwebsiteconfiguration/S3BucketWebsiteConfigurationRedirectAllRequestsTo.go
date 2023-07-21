package s3bucketwebsiteconfiguration


type S3BucketWebsiteConfigurationRedirectAllRequestsTo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/s3_bucket_website_configuration#host_name S3BucketWebsiteConfiguration#host_name}.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/s3_bucket_website_configuration#protocol S3BucketWebsiteConfiguration#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

