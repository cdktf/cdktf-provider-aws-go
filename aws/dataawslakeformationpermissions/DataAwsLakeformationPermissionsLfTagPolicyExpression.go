package dataawslakeformationpermissions


type DataAwsLakeformationPermissionsLfTagPolicyExpression struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/data-sources/lakeformation_permissions#key DataAwsLakeformationPermissions#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/data-sources/lakeformation_permissions#values DataAwsLakeformationPermissions#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

