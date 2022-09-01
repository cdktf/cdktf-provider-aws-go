package lex


type LexIntentFollowUpPrompt struct {
	// prompt block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#prompt LexIntent#prompt}
	Prompt *LexIntentFollowUpPromptPrompt `field:"required" json:"prompt" yaml:"prompt"`
	// rejection_statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#rejection_statement LexIntent#rejection_statement}
	RejectionStatement *LexIntentFollowUpPromptRejectionStatement `field:"required" json:"rejectionStatement" yaml:"rejectionStatement"`
}

