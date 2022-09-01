package cloudfront


type CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookies struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution#forward CloudfrontDistribution#forward}.
	Forward *string `field:"required" json:"forward" yaml:"forward"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution#whitelisted_names CloudfrontDistribution#whitelisted_names}.
	WhitelistedNames *[]*string `field:"optional" json:"whitelistedNames" yaml:"whitelistedNames"`
}

