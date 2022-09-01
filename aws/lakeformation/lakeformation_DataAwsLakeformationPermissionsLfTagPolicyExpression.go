package lakeformation


type DataAwsLakeformationPermissionsLfTagPolicyExpression struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#key DataAwsLakeformationPermissions#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/lakeformation_permissions#values DataAwsLakeformationPermissions#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

