// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2allowedimagessettings


type Ec2AllowedImagesSettingsImageCriterionCreationDateCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/ec2_allowed_images_settings#maximum_days_since_created Ec2AllowedImagesSettings#maximum_days_since_created}.
	MaximumDaysSinceCreated *float64 `field:"optional" json:"maximumDaysSinceCreated" yaml:"maximumDaysSinceCreated"`
}

