// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontresponseheaderspolicy


type CloudfrontResponseHeadersPolicySecurityHeadersConfig struct {
	// content_security_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/cloudfront_response_headers_policy#content_security_policy CloudfrontResponseHeadersPolicy#content_security_policy}
	ContentSecurityPolicy *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy `field:"optional" json:"contentSecurityPolicy" yaml:"contentSecurityPolicy"`
	// content_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/cloudfront_response_headers_policy#content_type_options CloudfrontResponseHeadersPolicy#content_type_options}
	ContentTypeOptions *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions `field:"optional" json:"contentTypeOptions" yaml:"contentTypeOptions"`
	// frame_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/cloudfront_response_headers_policy#frame_options CloudfrontResponseHeadersPolicy#frame_options}
	FrameOptions *CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions `field:"optional" json:"frameOptions" yaml:"frameOptions"`
	// referrer_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/cloudfront_response_headers_policy#referrer_policy CloudfrontResponseHeadersPolicy#referrer_policy}
	ReferrerPolicy *CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy `field:"optional" json:"referrerPolicy" yaml:"referrerPolicy"`
	// strict_transport_security block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/cloudfront_response_headers_policy#strict_transport_security CloudfrontResponseHeadersPolicy#strict_transport_security}
	StrictTransportSecurity *CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity `field:"optional" json:"strictTransportSecurity" yaml:"strictTransportSecurity"`
	// xss_protection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/cloudfront_response_headers_policy#xss_protection CloudfrontResponseHeadersPolicy#xss_protection}
	XssProtection *CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection `field:"optional" json:"xssProtection" yaml:"xssProtection"`
}

