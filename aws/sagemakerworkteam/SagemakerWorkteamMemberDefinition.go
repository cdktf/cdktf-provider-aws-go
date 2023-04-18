package sagemakerworkteam


type SagemakerWorkteamMemberDefinition struct {
	// cognito_member_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/sagemaker_workteam#cognito_member_definition SagemakerWorkteam#cognito_member_definition}
	CognitoMemberDefinition *SagemakerWorkteamMemberDefinitionCognitoMemberDefinition `field:"optional" json:"cognitoMemberDefinition" yaml:"cognitoMemberDefinition"`
	// oidc_member_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/sagemaker_workteam#oidc_member_definition SagemakerWorkteam#oidc_member_definition}
	OidcMemberDefinition *SagemakerWorkteamMemberDefinitionOidcMemberDefinition `field:"optional" json:"oidcMemberDefinition" yaml:"oidcMemberDefinition"`
}

