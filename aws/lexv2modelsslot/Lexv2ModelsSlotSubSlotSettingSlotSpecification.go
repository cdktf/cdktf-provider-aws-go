// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsslot


type Lexv2ModelsSlotSubSlotSettingSlotSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/lexv2models_slot#map_block_key Lexv2ModelsSlot#map_block_key}.
	MapBlockKey *string `field:"required" json:"mapBlockKey" yaml:"mapBlockKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/lexv2models_slot#slot_type_id Lexv2ModelsSlot#slot_type_id}.
	SlotTypeId *string `field:"required" json:"slotTypeId" yaml:"slotTypeId"`
	// value_elicitation_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/lexv2models_slot#value_elicitation_setting Lexv2ModelsSlot#value_elicitation_setting}
	ValueElicitationSetting interface{} `field:"optional" json:"valueElicitationSetting" yaml:"valueElicitationSetting"`
}

