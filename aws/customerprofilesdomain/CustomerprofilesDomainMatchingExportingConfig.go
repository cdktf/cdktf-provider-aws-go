// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customerprofilesdomain


type CustomerprofilesDomainMatchingExportingConfig struct {
	// s3_exporting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/customerprofiles_domain#s3_exporting CustomerprofilesDomain#s3_exporting}
	S3Exporting *CustomerprofilesDomainMatchingExportingConfigS3Exporting `field:"optional" json:"s3Exporting" yaml:"s3Exporting"`
}

