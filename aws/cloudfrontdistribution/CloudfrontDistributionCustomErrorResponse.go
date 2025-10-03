// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistribution


type CloudfrontDistributionCustomErrorResponse struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudfront_distribution#error_code CloudfrontDistribution#error_code}.
	ErrorCode *float64 `field:"required" json:"errorCode" yaml:"errorCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudfront_distribution#error_caching_min_ttl CloudfrontDistribution#error_caching_min_ttl}.
	ErrorCachingMinTtl *float64 `field:"optional" json:"errorCachingMinTtl" yaml:"errorCachingMinTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudfront_distribution#response_code CloudfrontDistribution#response_code}.
	ResponseCode *float64 `field:"optional" json:"responseCode" yaml:"responseCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudfront_distribution#response_page_path CloudfrontDistribution#response_page_path}.
	ResponsePagePath *string `field:"optional" json:"responsePagePath" yaml:"responsePagePath"`
}

