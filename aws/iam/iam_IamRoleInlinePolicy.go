package iam


type IamRoleInlinePolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iam_role#name IamRole#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/iam_role#policy IamRole#policy}.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

