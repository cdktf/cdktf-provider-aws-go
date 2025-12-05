// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2allowedimagessettings


type Ec2AllowedImagesSettingsImageCriterionDeprecationTimeCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ec2_allowed_images_settings#maximum_days_since_deprecated Ec2AllowedImagesSettings#maximum_days_since_deprecated}.
	MaximumDaysSinceDeprecated *float64 `field:"optional" json:"maximumDaysSinceDeprecated" yaml:"maximumDaysSinceDeprecated"`
}

