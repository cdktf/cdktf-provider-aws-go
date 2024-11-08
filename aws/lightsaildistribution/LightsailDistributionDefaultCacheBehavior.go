// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lightsaildistribution


type LightsailDistributionDefaultCacheBehavior struct {
	// The cache behavior of the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/lightsail_distribution#behavior LightsailDistribution#behavior}
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
}

