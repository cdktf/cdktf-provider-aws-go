// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawslicensemanagergrants


type DataAwsLicensemanagerGrantsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/data-sources/licensemanager_grants#name DataAwsLicensemanagerGrants#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/data-sources/licensemanager_grants#values DataAwsLicensemanagerGrants#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

