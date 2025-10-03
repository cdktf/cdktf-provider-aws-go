// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontcachepolicy


type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudfront_cache_policy#query_string_behavior CloudfrontCachePolicy#query_string_behavior}.
	QueryStringBehavior *string `field:"required" json:"queryStringBehavior" yaml:"queryStringBehavior"`
	// query_strings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudfront_cache_policy#query_strings CloudfrontCachePolicy#query_strings}
	QueryStrings *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings `field:"optional" json:"queryStrings" yaml:"queryStrings"`
}

