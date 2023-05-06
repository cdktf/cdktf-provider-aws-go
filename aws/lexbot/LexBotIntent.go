package lexbot


type LexBotIntent struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/lex_bot#intent_name LexBot#intent_name}.
	IntentName *string `field:"required" json:"intentName" yaml:"intentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/lex_bot#intent_version LexBot#intent_version}.
	IntentVersion *string `field:"required" json:"intentVersion" yaml:"intentVersion"`
}

