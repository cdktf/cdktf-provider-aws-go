package cloudfrontdistribution


type CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/cloudfront_distribution#forward CloudfrontDistribution#forward}.
	Forward *string `field:"required" json:"forward" yaml:"forward"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/cloudfront_distribution#whitelisted_names CloudfrontDistribution#whitelisted_names}.
	WhitelistedNames *[]*string `field:"optional" json:"whitelistedNames" yaml:"whitelistedNames"`
}

