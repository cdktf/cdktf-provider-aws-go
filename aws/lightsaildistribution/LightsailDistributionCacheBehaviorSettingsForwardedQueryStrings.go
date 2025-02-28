// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lightsaildistribution


type LightsailDistributionCacheBehaviorSettingsForwardedQueryStrings struct {
	// Indicates whether the distribution forwards and caches based on query strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/lightsail_distribution#option LightsailDistribution#option}
	Option interface{} `field:"optional" json:"option" yaml:"option"`
	// The specific query strings that the distribution forwards to the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/lightsail_distribution#query_strings_allowed_list LightsailDistribution#query_strings_allowed_list}
	QueryStringsAllowedList *[]*string `field:"optional" json:"queryStringsAllowedList" yaml:"queryStringsAllowedList"`
}

