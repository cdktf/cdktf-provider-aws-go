// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customerprofilesdomain


type CustomerprofilesDomainMatching struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/customerprofiles_domain#enabled CustomerprofilesDomain#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// auto_merging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/customerprofiles_domain#auto_merging CustomerprofilesDomain#auto_merging}
	AutoMerging *CustomerprofilesDomainMatchingAutoMerging `field:"optional" json:"autoMerging" yaml:"autoMerging"`
	// exporting_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/customerprofiles_domain#exporting_config CustomerprofilesDomain#exporting_config}
	ExportingConfig *CustomerprofilesDomainMatchingExportingConfig `field:"optional" json:"exportingConfig" yaml:"exportingConfig"`
	// job_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/customerprofiles_domain#job_schedule CustomerprofilesDomain#job_schedule}
	JobSchedule *CustomerprofilesDomainMatchingJobSchedule `field:"optional" json:"jobSchedule" yaml:"jobSchedule"`
}

