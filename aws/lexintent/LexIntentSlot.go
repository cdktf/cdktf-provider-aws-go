// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexintent


type LexIntentSlot struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lex_intent#name LexIntent#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lex_intent#slot_constraint LexIntent#slot_constraint}.
	SlotConstraint *string `field:"required" json:"slotConstraint" yaml:"slotConstraint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lex_intent#slot_type LexIntent#slot_type}.
	SlotType *string `field:"required" json:"slotType" yaml:"slotType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lex_intent#description LexIntent#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lex_intent#priority LexIntent#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lex_intent#response_card LexIntent#response_card}.
	ResponseCard *string `field:"optional" json:"responseCard" yaml:"responseCard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lex_intent#sample_utterances LexIntent#sample_utterances}.
	SampleUtterances *[]*string `field:"optional" json:"sampleUtterances" yaml:"sampleUtterances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lex_intent#slot_type_version LexIntent#slot_type_version}.
	SlotTypeVersion *string `field:"optional" json:"slotTypeVersion" yaml:"slotTypeVersion"`
	// value_elicitation_prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lex_intent#value_elicitation_prompt LexIntent#value_elicitation_prompt}
	ValueElicitationPrompt *LexIntentSlotValueElicitationPrompt `field:"optional" json:"valueElicitationPrompt" yaml:"valueElicitationPrompt"`
}

