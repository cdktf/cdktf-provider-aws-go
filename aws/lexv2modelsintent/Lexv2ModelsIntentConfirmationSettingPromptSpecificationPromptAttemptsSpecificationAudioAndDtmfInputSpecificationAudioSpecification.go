// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecificationAudioSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/lexv2models_intent#end_timeout_ms Lexv2ModelsIntent#end_timeout_ms}.
	EndTimeoutMs *float64 `field:"required" json:"endTimeoutMs" yaml:"endTimeoutMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/lexv2models_intent#max_length_ms Lexv2ModelsIntent#max_length_ms}.
	MaxLengthMs *float64 `field:"required" json:"maxLengthMs" yaml:"maxLengthMs"`
}

