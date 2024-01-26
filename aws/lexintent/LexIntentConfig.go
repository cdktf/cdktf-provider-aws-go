// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexintent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LexIntentConfig struct {
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
	// fulfillment_activity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/lex_intent#fulfillment_activity LexIntent#fulfillment_activity}
	FulfillmentActivity *LexIntentFulfillmentActivity `field:"required" json:"fulfillmentActivity" yaml:"fulfillmentActivity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/lex_intent#name LexIntent#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// conclusion_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/lex_intent#conclusion_statement LexIntent#conclusion_statement}
	ConclusionStatement *LexIntentConclusionStatement `field:"optional" json:"conclusionStatement" yaml:"conclusionStatement"`
	// confirmation_prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/lex_intent#confirmation_prompt LexIntent#confirmation_prompt}
	ConfirmationPrompt *LexIntentConfirmationPrompt `field:"optional" json:"confirmationPrompt" yaml:"confirmationPrompt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/lex_intent#create_version LexIntent#create_version}.
	CreateVersion interface{} `field:"optional" json:"createVersion" yaml:"createVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/lex_intent#description LexIntent#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dialog_code_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/lex_intent#dialog_code_hook LexIntent#dialog_code_hook}
	DialogCodeHook *LexIntentDialogCodeHook `field:"optional" json:"dialogCodeHook" yaml:"dialogCodeHook"`
	// follow_up_prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/lex_intent#follow_up_prompt LexIntent#follow_up_prompt}
	FollowUpPrompt *LexIntentFollowUpPrompt `field:"optional" json:"followUpPrompt" yaml:"followUpPrompt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/lex_intent#id LexIntent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/lex_intent#parent_intent_signature LexIntent#parent_intent_signature}.
	ParentIntentSignature *string `field:"optional" json:"parentIntentSignature" yaml:"parentIntentSignature"`
	// rejection_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/lex_intent#rejection_statement LexIntent#rejection_statement}
	RejectionStatement *LexIntentRejectionStatement `field:"optional" json:"rejectionStatement" yaml:"rejectionStatement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/lex_intent#sample_utterances LexIntent#sample_utterances}.
	SampleUtterances *[]*string `field:"optional" json:"sampleUtterances" yaml:"sampleUtterances"`
	// slot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/lex_intent#slot LexIntent#slot}
	Slot interface{} `field:"optional" json:"slot" yaml:"slot"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/lex_intent#timeouts LexIntent#timeouts}
	Timeouts *LexIntentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

