package kendraindex


type KendraIndexUserTokenConfigurations struct {
	// json_token_type_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/kendra_index#json_token_type_configuration KendraIndex#json_token_type_configuration}
	JsonTokenTypeConfiguration *KendraIndexUserTokenConfigurationsJsonTokenTypeConfiguration `field:"optional" json:"jsonTokenTypeConfiguration" yaml:"jsonTokenTypeConfiguration"`
	// jwt_token_type_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/kendra_index#jwt_token_type_configuration KendraIndex#jwt_token_type_configuration}
	JwtTokenTypeConfiguration *KendraIndexUserTokenConfigurationsJwtTokenTypeConfiguration `field:"optional" json:"jwtTokenTypeConfiguration" yaml:"jwtTokenTypeConfiguration"`
}

