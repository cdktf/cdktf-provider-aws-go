// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customerprofilesdomain


type CustomerprofilesDomainRuleBasedMatchingExportingConfig struct {
	// s3_exporting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/customerprofiles_domain#s3_exporting CustomerprofilesDomain#s3_exporting}
	S3Exporting *CustomerprofilesDomainRuleBasedMatchingExportingConfigS3Exporting `field:"optional" json:"s3Exporting" yaml:"s3Exporting"`
}

