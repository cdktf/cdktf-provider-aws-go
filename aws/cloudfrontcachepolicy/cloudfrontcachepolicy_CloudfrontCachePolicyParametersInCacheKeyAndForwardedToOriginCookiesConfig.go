package cloudfrontcachepolicy


type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy#cookie_behavior CloudfrontCachePolicy#cookie_behavior}.
	CookieBehavior *string `field:"required" json:"cookieBehavior" yaml:"cookieBehavior"`
	// cookies block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy#cookies CloudfrontCachePolicy#cookies}
	Cookies *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies `field:"optional" json:"cookies" yaml:"cookies"`
}

