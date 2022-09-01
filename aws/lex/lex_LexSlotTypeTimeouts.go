package lex


type LexSlotTypeTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type#create LexSlotType#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type#delete LexSlotType#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type#update LexSlotType#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

