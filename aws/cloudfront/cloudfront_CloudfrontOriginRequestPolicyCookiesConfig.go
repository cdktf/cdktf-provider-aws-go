package cloudfront


type CloudfrontOriginRequestPolicyCookiesConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy#cookie_behavior CloudfrontOriginRequestPolicy#cookie_behavior}.
	CookieBehavior *string `field:"required" json:"cookieBehavior" yaml:"cookieBehavior"`
	// cookies block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy#cookies CloudfrontOriginRequestPolicy#cookies}
	Cookies *CloudfrontOriginRequestPolicyCookiesConfigCookies `field:"optional" json:"cookies" yaml:"cookies"`
}

