// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontoriginrequestpolicy


type CloudfrontOriginRequestPolicyCookiesConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/cloudfront_origin_request_policy#cookie_behavior CloudfrontOriginRequestPolicy#cookie_behavior}.
	CookieBehavior *string `field:"required" json:"cookieBehavior" yaml:"cookieBehavior"`
	// cookies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/cloudfront_origin_request_policy#cookies CloudfrontOriginRequestPolicy#cookies}
	Cookies *CloudfrontOriginRequestPolicyCookiesConfigCookies `field:"optional" json:"cookies" yaml:"cookies"`
}

