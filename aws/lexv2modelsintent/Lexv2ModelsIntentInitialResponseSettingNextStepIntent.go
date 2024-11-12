// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentInitialResponseSettingNextStepIntent struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/lexv2models_intent#name Lexv2ModelsIntent#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// slot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/lexv2models_intent#slot Lexv2ModelsIntent#slot}
	Slot interface{} `field:"optional" json:"slot" yaml:"slot"`
}

