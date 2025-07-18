// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontoriginrequestpolicy


type CloudfrontOriginRequestPolicyQueryStringsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/cloudfront_origin_request_policy#query_string_behavior CloudfrontOriginRequestPolicy#query_string_behavior}.
	QueryStringBehavior *string `field:"required" json:"queryStringBehavior" yaml:"queryStringBehavior"`
	// query_strings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/cloudfront_origin_request_policy#query_strings CloudfrontOriginRequestPolicy#query_strings}
	QueryStrings *CloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings `field:"optional" json:"queryStrings" yaml:"queryStrings"`
}

