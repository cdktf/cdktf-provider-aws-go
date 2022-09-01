// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KendraIndexUserTokenConfigurationsJwtTokenTypeConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#key_location KendraIndex#key_location}.
	KeyLocation *string `field:"required" json:"keyLocation" yaml:"keyLocation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#claim_regex KendraIndex#claim_regex}.
	ClaimRegex *string `field:"optional" json:"claimRegex" yaml:"claimRegex"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#group_attribute_field KendraIndex#group_attribute_field}.
	GroupAttributeField *string `field:"optional" json:"groupAttributeField" yaml:"groupAttributeField"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#issuer KendraIndex#issuer}.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#secrets_manager_arn KendraIndex#secrets_manager_arn}.
	SecretsManagerArn *string `field:"optional" json:"secretsManagerArn" yaml:"secretsManagerArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#url KendraIndex#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#user_name_attribute_field KendraIndex#user_name_attribute_field}.
	UserNameAttributeField *string `field:"optional" json:"userNameAttributeField" yaml:"userNameAttributeField"`
}

