package dataawslicensemanagergrants


type DataAwsLicensemanagerGrantsFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/licensemanager_grants#name DataAwsLicensemanagerGrants#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/licensemanager_grants#values DataAwsLicensemanagerGrants#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

