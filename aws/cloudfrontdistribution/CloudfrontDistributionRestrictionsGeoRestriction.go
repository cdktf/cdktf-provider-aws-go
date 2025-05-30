// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistribution


type CloudfrontDistributionRestrictionsGeoRestriction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/cloudfront_distribution#restriction_type CloudfrontDistribution#restriction_type}.
	RestrictionType *string `field:"required" json:"restrictionType" yaml:"restrictionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/cloudfront_distribution#locations CloudfrontDistribution#locations}.
	Locations *[]*string `field:"optional" json:"locations" yaml:"locations"`
}

