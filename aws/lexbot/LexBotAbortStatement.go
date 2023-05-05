package lexbot


type LexBotAbortStatement struct {
	// message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/lex_bot#message LexBot#message}
	Message interface{} `field:"required" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/lex_bot#response_card LexBot#response_card}.
	ResponseCard *string `field:"optional" json:"responseCard" yaml:"responseCard"`
}

