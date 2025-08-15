// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentSlotPriority struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/lexv2models_intent#priority Lexv2ModelsIntent#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/lexv2models_intent#slot_id Lexv2ModelsIntent#slot_id}.
	SlotId *string `field:"required" json:"slotId" yaml:"slotId"`
}

