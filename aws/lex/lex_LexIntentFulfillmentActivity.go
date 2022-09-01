package lex


type LexIntentFulfillmentActivity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#type LexIntent#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// code_hook block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#code_hook LexIntent#code_hook}
	CodeHook *LexIntentFulfillmentActivityCodeHook `field:"optional" json:"codeHook" yaml:"codeHook"`
}

