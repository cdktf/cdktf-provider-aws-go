package lex


type LexSlotTypeEnumerationValue struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type#value LexSlotType#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_slot_type#synonyms LexSlotType#synonyms}.
	Synonyms *[]*string `field:"optional" json:"synonyms" yaml:"synonyms"`
}

