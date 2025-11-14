// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Lexv2ModelsIntentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lexv2models_intent#bot_id Lexv2ModelsIntent#bot_id}.
	BotId *string `field:"required" json:"botId" yaml:"botId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lexv2models_intent#bot_version Lexv2ModelsIntent#bot_version}.
	BotVersion *string `field:"required" json:"botVersion" yaml:"botVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lexv2models_intent#locale_id Lexv2ModelsIntent#locale_id}.
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lexv2models_intent#name Lexv2ModelsIntent#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// closing_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lexv2models_intent#closing_setting Lexv2ModelsIntent#closing_setting}
	ClosingSetting interface{} `field:"optional" json:"closingSetting" yaml:"closingSetting"`
	// confirmation_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lexv2models_intent#confirmation_setting Lexv2ModelsIntent#confirmation_setting}
	ConfirmationSetting interface{} `field:"optional" json:"confirmationSetting" yaml:"confirmationSetting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lexv2models_intent#description Lexv2ModelsIntent#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dialog_code_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lexv2models_intent#dialog_code_hook Lexv2ModelsIntent#dialog_code_hook}
	DialogCodeHook interface{} `field:"optional" json:"dialogCodeHook" yaml:"dialogCodeHook"`
	// fulfillment_code_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lexv2models_intent#fulfillment_code_hook Lexv2ModelsIntent#fulfillment_code_hook}
	FulfillmentCodeHook interface{} `field:"optional" json:"fulfillmentCodeHook" yaml:"fulfillmentCodeHook"`
	// initial_response_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lexv2models_intent#initial_response_setting Lexv2ModelsIntent#initial_response_setting}
	InitialResponseSetting interface{} `field:"optional" json:"initialResponseSetting" yaml:"initialResponseSetting"`
	// input_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lexv2models_intent#input_context Lexv2ModelsIntent#input_context}
	InputContext interface{} `field:"optional" json:"inputContext" yaml:"inputContext"`
	// kendra_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lexv2models_intent#kendra_configuration Lexv2ModelsIntent#kendra_configuration}
	KendraConfiguration interface{} `field:"optional" json:"kendraConfiguration" yaml:"kendraConfiguration"`
	// output_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lexv2models_intent#output_context Lexv2ModelsIntent#output_context}
	OutputContext interface{} `field:"optional" json:"outputContext" yaml:"outputContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lexv2models_intent#parent_intent_signature Lexv2ModelsIntent#parent_intent_signature}.
	ParentIntentSignature *string `field:"optional" json:"parentIntentSignature" yaml:"parentIntentSignature"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lexv2models_intent#region Lexv2ModelsIntent#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// sample_utterance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lexv2models_intent#sample_utterance Lexv2ModelsIntent#sample_utterance}
	SampleUtterance interface{} `field:"optional" json:"sampleUtterance" yaml:"sampleUtterance"`
	// slot_priority block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lexv2models_intent#slot_priority Lexv2ModelsIntent#slot_priority}
	SlotPriority interface{} `field:"optional" json:"slotPriority" yaml:"slotPriority"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/lexv2models_intent#timeouts Lexv2ModelsIntent#timeouts}
	Timeouts *Lexv2ModelsIntentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

