// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customerprofilesdomain


type CustomerprofilesDomainMatchingJobSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/customerprofiles_domain#day_of_the_week CustomerprofilesDomain#day_of_the_week}.
	DayOfTheWeek *string `field:"required" json:"dayOfTheWeek" yaml:"dayOfTheWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/customerprofiles_domain#time CustomerprofilesDomain#time}.
	Time *string `field:"required" json:"time" yaml:"time"`
}

