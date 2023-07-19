package lexslottype


type LexSlotTypeEnumerationValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/lex_slot_type#value LexSlotType#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/lex_slot_type#synonyms LexSlotType#synonyms}.
	Synonyms *[]*string `field:"optional" json:"synonyms" yaml:"synonyms"`
}

