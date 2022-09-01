package lex


type LexIntentDialogCodeHook struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#message_version LexIntent#message_version}.
	MessageVersion *string `field:"required" json:"messageVersion" yaml:"messageVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_intent#uri LexIntent#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

