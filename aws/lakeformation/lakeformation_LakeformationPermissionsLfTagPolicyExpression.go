package lakeformation


type LakeformationPermissionsLfTagPolicyExpression struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#key LakeformationPermissions#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lakeformation_permissions#values LakeformationPermissions#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

