// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsslot


type Lexv2ModelsSlotSubSlotSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/lexv2models_slot#expression Lexv2ModelsSlot#expression}.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// slot_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/lexv2models_slot#slot_specification Lexv2ModelsSlot#slot_specification}
	SlotSpecification interface{} `field:"optional" json:"slotSpecification" yaml:"slotSpecification"`
}

