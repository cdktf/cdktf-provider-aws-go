package lexintent


type LexIntentSlotValueElicitationPrompt struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lex_intent#max_attempts LexIntent#max_attempts}.
	MaxAttempts *float64 `field:"required" json:"maxAttempts" yaml:"maxAttempts"`
	// message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lex_intent#message LexIntent#message}
	Message interface{} `field:"required" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lex_intent#response_card LexIntent#response_card}.
	ResponseCard *string `field:"optional" json:"responseCard" yaml:"responseCard"`
}

