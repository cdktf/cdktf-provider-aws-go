// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KendraIndexUserTokenConfigurationsJsonTokenTypeConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#group_attribute_field KendraIndex#group_attribute_field}.
	GroupAttributeField *string `field:"required" json:"groupAttributeField" yaml:"groupAttributeField"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#user_name_attribute_field KendraIndex#user_name_attribute_field}.
	UserNameAttributeField *string `field:"required" json:"userNameAttributeField" yaml:"userNameAttributeField"`
}

