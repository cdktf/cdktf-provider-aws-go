// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsslottype

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Lexv2ModelsSlotTypeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lexv2models_slot_type#bot_id Lexv2ModelsSlotType#bot_id}.
	BotId *string `field:"required" json:"botId" yaml:"botId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lexv2models_slot_type#bot_version Lexv2ModelsSlotType#bot_version}.
	BotVersion *string `field:"required" json:"botVersion" yaml:"botVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lexv2models_slot_type#locale_id Lexv2ModelsSlotType#locale_id}.
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lexv2models_slot_type#name Lexv2ModelsSlotType#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// composite_slot_type_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lexv2models_slot_type#composite_slot_type_setting Lexv2ModelsSlotType#composite_slot_type_setting}
	CompositeSlotTypeSetting interface{} `field:"optional" json:"compositeSlotTypeSetting" yaml:"compositeSlotTypeSetting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lexv2models_slot_type#description Lexv2ModelsSlotType#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// external_source_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lexv2models_slot_type#external_source_setting Lexv2ModelsSlotType#external_source_setting}
	ExternalSourceSetting interface{} `field:"optional" json:"externalSourceSetting" yaml:"externalSourceSetting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lexv2models_slot_type#parent_slot_type_signature Lexv2ModelsSlotType#parent_slot_type_signature}.
	ParentSlotTypeSignature *string `field:"optional" json:"parentSlotTypeSignature" yaml:"parentSlotTypeSignature"`
	// slot_type_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lexv2models_slot_type#slot_type_values Lexv2ModelsSlotType#slot_type_values}
	SlotTypeValues interface{} `field:"optional" json:"slotTypeValues" yaml:"slotTypeValues"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lexv2models_slot_type#timeouts Lexv2ModelsSlotType#timeouts}
	Timeouts *Lexv2ModelsSlotTypeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// value_selection_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lexv2models_slot_type#value_selection_setting Lexv2ModelsSlotType#value_selection_setting}
	ValueSelectionSetting interface{} `field:"optional" json:"valueSelectionSetting" yaml:"valueSelectionSetting"`
}

