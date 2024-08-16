// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsslot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Lexv2ModelsSlotConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/lexv2models_slot#bot_id Lexv2ModelsSlot#bot_id}.
	BotId *string `field:"required" json:"botId" yaml:"botId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/lexv2models_slot#bot_version Lexv2ModelsSlot#bot_version}.
	BotVersion *string `field:"required" json:"botVersion" yaml:"botVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/lexv2models_slot#intent_id Lexv2ModelsSlot#intent_id}.
	IntentId *string `field:"required" json:"intentId" yaml:"intentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/lexv2models_slot#locale_id Lexv2ModelsSlot#locale_id}.
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/lexv2models_slot#name Lexv2ModelsSlot#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/lexv2models_slot#description Lexv2ModelsSlot#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// multiple_values_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/lexv2models_slot#multiple_values_setting Lexv2ModelsSlot#multiple_values_setting}
	MultipleValuesSetting interface{} `field:"optional" json:"multipleValuesSetting" yaml:"multipleValuesSetting"`
	// obfuscation_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/lexv2models_slot#obfuscation_setting Lexv2ModelsSlot#obfuscation_setting}
	ObfuscationSetting interface{} `field:"optional" json:"obfuscationSetting" yaml:"obfuscationSetting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/lexv2models_slot#slot_type_id Lexv2ModelsSlot#slot_type_id}.
	SlotTypeId *string `field:"optional" json:"slotTypeId" yaml:"slotTypeId"`
	// sub_slot_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/lexv2models_slot#sub_slot_setting Lexv2ModelsSlot#sub_slot_setting}
	SubSlotSetting interface{} `field:"optional" json:"subSlotSetting" yaml:"subSlotSetting"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/lexv2models_slot#timeouts Lexv2ModelsSlot#timeouts}
	Timeouts *Lexv2ModelsSlotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// value_elicitation_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/lexv2models_slot#value_elicitation_setting Lexv2ModelsSlot#value_elicitation_setting}
	ValueElicitationSetting interface{} `field:"optional" json:"valueElicitationSetting" yaml:"valueElicitationSetting"`
}

