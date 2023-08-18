package lakeformationpermissions


type LakeformationPermissionsLfTagPolicyExpression struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/lakeformation_permissions#key LakeformationPermissions#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/lakeformation_permissions#values LakeformationPermissions#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

