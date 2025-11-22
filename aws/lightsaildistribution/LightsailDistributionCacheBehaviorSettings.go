// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lightsaildistribution


type LightsailDistributionCacheBehaviorSettings struct {
	// The HTTP methods that are processed and forwarded to the distribution's origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/lightsail_distribution#allowed_http_methods LightsailDistribution#allowed_http_methods}
	AllowedHttpMethods *string `field:"optional" json:"allowedHttpMethods" yaml:"allowedHttpMethods"`
	// The HTTP method responses that are cached by your distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/lightsail_distribution#cached_http_methods LightsailDistribution#cached_http_methods}
	CachedHttpMethods *string `field:"optional" json:"cachedHttpMethods" yaml:"cachedHttpMethods"`
	// The default amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the content has been updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/lightsail_distribution#default_ttl LightsailDistribution#default_ttl}
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// forwarded_cookies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/lightsail_distribution#forwarded_cookies LightsailDistribution#forwarded_cookies}
	ForwardedCookies *LightsailDistributionCacheBehaviorSettingsForwardedCookies `field:"optional" json:"forwardedCookies" yaml:"forwardedCookies"`
	// forwarded_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/lightsail_distribution#forwarded_headers LightsailDistribution#forwarded_headers}
	ForwardedHeaders *LightsailDistributionCacheBehaviorSettingsForwardedHeaders `field:"optional" json:"forwardedHeaders" yaml:"forwardedHeaders"`
	// forwarded_query_strings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/lightsail_distribution#forwarded_query_strings LightsailDistribution#forwarded_query_strings}
	ForwardedQueryStrings *LightsailDistributionCacheBehaviorSettingsForwardedQueryStrings `field:"optional" json:"forwardedQueryStrings" yaml:"forwardedQueryStrings"`
	// The maximum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/lightsail_distribution#maximum_ttl LightsailDistribution#maximum_ttl}
	MaximumTtl *float64 `field:"optional" json:"maximumTtl" yaml:"maximumTtl"`
	// The minimum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/lightsail_distribution#minimum_ttl LightsailDistribution#minimum_ttl}
	MinimumTtl *float64 `field:"optional" json:"minimumTtl" yaml:"minimumTtl"`
}

