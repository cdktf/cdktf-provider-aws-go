package cognito


type CognitoUserPoolSchemaNumberAttributeConstraints struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#max_value CognitoUserPool#max_value}.
	MaxValue *string `field:"optional" json:"maxValue" yaml:"maxValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#min_value CognitoUserPool#min_value}.
	MinValue *string `field:"optional" json:"minValue" yaml:"minValue"`
}

