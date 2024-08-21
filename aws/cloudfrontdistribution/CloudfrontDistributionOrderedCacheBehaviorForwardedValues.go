// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistribution


type CloudfrontDistributionOrderedCacheBehaviorForwardedValues struct {
	// cookies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/cloudfront_distribution#cookies CloudfrontDistribution#cookies}
	Cookies *CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookies `field:"required" json:"cookies" yaml:"cookies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/cloudfront_distribution#query_string CloudfrontDistribution#query_string}.
	QueryString interface{} `field:"required" json:"queryString" yaml:"queryString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/cloudfront_distribution#headers CloudfrontDistribution#headers}.
	Headers *[]*string `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/cloudfront_distribution#query_string_cache_keys CloudfrontDistribution#query_string_cache_keys}.
	QueryStringCacheKeys *[]*string `field:"optional" json:"queryStringCacheKeys" yaml:"queryStringCacheKeys"`
}

