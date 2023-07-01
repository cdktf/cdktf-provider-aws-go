package iamrole


type IamRoleInlinePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/iam_role#name IamRole#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/iam_role#policy IamRole#policy}.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

