package cloudfront


type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy#header_behavior CloudfrontCachePolicy#header_behavior}.
	HeaderBehavior *string `field:"optional" json:"headerBehavior" yaml:"headerBehavior"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy#headers CloudfrontCachePolicy#headers}
	Headers *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders `field:"optional" json:"headers" yaml:"headers"`
}

