// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontcachepolicy


type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/cloudfront_cache_policy#header_behavior CloudfrontCachePolicy#header_behavior}.
	HeaderBehavior *string `field:"optional" json:"headerBehavior" yaml:"headerBehavior"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/cloudfront_cache_policy#headers CloudfrontCachePolicy#headers}
	Headers *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders `field:"optional" json:"headers" yaml:"headers"`
}

