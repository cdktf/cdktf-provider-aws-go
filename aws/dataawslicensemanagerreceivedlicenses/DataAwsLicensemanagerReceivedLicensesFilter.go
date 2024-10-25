// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawslicensemanagerreceivedlicenses


type DataAwsLicensemanagerReceivedLicensesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/data-sources/licensemanager_received_licenses#name DataAwsLicensemanagerReceivedLicenses#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/data-sources/licensemanager_received_licenses#values DataAwsLicensemanagerReceivedLicenses#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

