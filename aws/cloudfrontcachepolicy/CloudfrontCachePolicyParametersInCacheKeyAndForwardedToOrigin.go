// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontcachepolicy


type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin struct {
	// cookies_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/cloudfront_cache_policy#cookies_config CloudfrontCachePolicy#cookies_config}
	CookiesConfig *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig `field:"required" json:"cookiesConfig" yaml:"cookiesConfig"`
	// headers_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/cloudfront_cache_policy#headers_config CloudfrontCachePolicy#headers_config}
	HeadersConfig *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig `field:"required" json:"headersConfig" yaml:"headersConfig"`
	// query_strings_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/cloudfront_cache_policy#query_strings_config CloudfrontCachePolicy#query_strings_config}
	QueryStringsConfig *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig `field:"required" json:"queryStringsConfig" yaml:"queryStringsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/cloudfront_cache_policy#enable_accept_encoding_brotli CloudfrontCachePolicy#enable_accept_encoding_brotli}.
	EnableAcceptEncodingBrotli interface{} `field:"optional" json:"enableAcceptEncodingBrotli" yaml:"enableAcceptEncodingBrotli"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/cloudfront_cache_policy#enable_accept_encoding_gzip CloudfrontCachePolicy#enable_accept_encoding_gzip}.
	EnableAcceptEncodingGzip interface{} `field:"optional" json:"enableAcceptEncodingGzip" yaml:"enableAcceptEncodingGzip"`
}

