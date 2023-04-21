package connectbotassociation


type ConnectBotAssociationLexBot struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/connect_bot_association#name ConnectBotAssociation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/connect_bot_association#lex_region ConnectBotAssociation#lex_region}.
	LexRegion *string `field:"optional" json:"lexRegion" yaml:"lexRegion"`
}

