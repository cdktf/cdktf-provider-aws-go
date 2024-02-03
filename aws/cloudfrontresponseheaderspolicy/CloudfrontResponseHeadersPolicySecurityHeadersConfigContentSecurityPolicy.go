// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontresponseheaderspolicy


type CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cloudfront_response_headers_policy#content_security_policy CloudfrontResponseHeadersPolicy#content_security_policy}.
	ContentSecurityPolicy *string `field:"required" json:"contentSecurityPolicy" yaml:"contentSecurityPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cloudfront_response_headers_policy#override CloudfrontResponseHeadersPolicy#override}.
	Override interface{} `field:"required" json:"override" yaml:"override"`
}

