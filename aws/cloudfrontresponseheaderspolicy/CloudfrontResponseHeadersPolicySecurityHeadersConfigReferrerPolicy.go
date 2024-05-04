// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontresponseheaderspolicy


type CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/cloudfront_response_headers_policy#override CloudfrontResponseHeadersPolicy#override}.
	Override interface{} `field:"required" json:"override" yaml:"override"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/cloudfront_response_headers_policy#referrer_policy CloudfrontResponseHeadersPolicy#referrer_policy}.
	ReferrerPolicy *string `field:"required" json:"referrerPolicy" yaml:"referrerPolicy"`
}

