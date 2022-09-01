package cloudfront


type CloudfrontDistributionLoggingConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution#bucket CloudfrontDistribution#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution#include_cookies CloudfrontDistribution#include_cookies}.
	IncludeCookies interface{} `field:"optional" json:"includeCookies" yaml:"includeCookies"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution#prefix CloudfrontDistribution#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

