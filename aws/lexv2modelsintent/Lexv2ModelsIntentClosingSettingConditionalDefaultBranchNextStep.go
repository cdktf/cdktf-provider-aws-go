// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentClosingSettingConditionalDefaultBranchNextStep struct {
	// dialog_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.77.0/docs/resources/lexv2models_intent#dialog_action Lexv2ModelsIntent#dialog_action}
	DialogAction interface{} `field:"optional" json:"dialogAction" yaml:"dialogAction"`
	// intent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.77.0/docs/resources/lexv2models_intent#intent Lexv2ModelsIntent#intent}
	Intent interface{} `field:"optional" json:"intent" yaml:"intent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.77.0/docs/resources/lexv2models_intent#session_attributes Lexv2ModelsIntent#session_attributes}.
	SessionAttributes *map[string]*string `field:"optional" json:"sessionAttributes" yaml:"sessionAttributes"`
}

