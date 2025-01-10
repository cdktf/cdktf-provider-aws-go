// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentConfirmationSettingElicitationCodeHook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/lexv2models_intent#enable_code_hook_invocation Lexv2ModelsIntent#enable_code_hook_invocation}.
	EnableCodeHookInvocation interface{} `field:"optional" json:"enableCodeHookInvocation" yaml:"enableCodeHookInvocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/lexv2models_intent#invocation_label Lexv2ModelsIntent#invocation_label}.
	InvocationLabel *string `field:"optional" json:"invocationLabel" yaml:"invocationLabel"`
}

