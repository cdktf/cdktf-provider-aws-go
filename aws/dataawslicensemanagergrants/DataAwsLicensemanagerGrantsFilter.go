package dataawslicensemanagergrants


type DataAwsLicensemanagerGrantsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/data-sources/licensemanager_grants#name DataAwsLicensemanagerGrants#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/data-sources/licensemanager_grants#values DataAwsLicensemanagerGrants#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

