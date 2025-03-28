// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontresponseheaderspolicy


type CloudfrontResponseHeadersPolicyCorsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_credentials CloudfrontResponseHeadersPolicy#access_control_allow_credentials}.
	AccessControlAllowCredentials interface{} `field:"required" json:"accessControlAllowCredentials" yaml:"accessControlAllowCredentials"`
	// access_control_allow_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_headers CloudfrontResponseHeadersPolicy#access_control_allow_headers}
	AccessControlAllowHeaders *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders `field:"required" json:"accessControlAllowHeaders" yaml:"accessControlAllowHeaders"`
	// access_control_allow_methods block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_methods CloudfrontResponseHeadersPolicy#access_control_allow_methods}
	AccessControlAllowMethods *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods `field:"required" json:"accessControlAllowMethods" yaml:"accessControlAllowMethods"`
	// access_control_allow_origins block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_origins CloudfrontResponseHeadersPolicy#access_control_allow_origins}
	AccessControlAllowOrigins *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins `field:"required" json:"accessControlAllowOrigins" yaml:"accessControlAllowOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/cloudfront_response_headers_policy#origin_override CloudfrontResponseHeadersPolicy#origin_override}.
	OriginOverride interface{} `field:"required" json:"originOverride" yaml:"originOverride"`
	// access_control_expose_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/cloudfront_response_headers_policy#access_control_expose_headers CloudfrontResponseHeadersPolicy#access_control_expose_headers}
	AccessControlExposeHeaders *CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders `field:"optional" json:"accessControlExposeHeaders" yaml:"accessControlExposeHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/cloudfront_response_headers_policy#access_control_max_age_sec CloudfrontResponseHeadersPolicy#access_control_max_age_sec}.
	AccessControlMaxAgeSec *float64 `field:"optional" json:"accessControlMaxAgeSec" yaml:"accessControlMaxAgeSec"`
}

