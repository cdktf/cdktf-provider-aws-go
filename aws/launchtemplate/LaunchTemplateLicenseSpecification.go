// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package launchtemplate


type LaunchTemplateLicenseSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/launch_template#license_configuration_arn LaunchTemplate#license_configuration_arn}.
	LicenseConfigurationArn *string `field:"required" json:"licenseConfigurationArn" yaml:"licenseConfigurationArn"`
}

