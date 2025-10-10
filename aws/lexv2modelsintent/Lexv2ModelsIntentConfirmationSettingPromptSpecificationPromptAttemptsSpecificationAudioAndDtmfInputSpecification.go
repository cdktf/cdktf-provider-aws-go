// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/lexv2models_intent#start_timeout_ms Lexv2ModelsIntent#start_timeout_ms}.
	StartTimeoutMs *float64 `field:"required" json:"startTimeoutMs" yaml:"startTimeoutMs"`
	// audio_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/lexv2models_intent#audio_specification Lexv2ModelsIntent#audio_specification}
	AudioSpecification interface{} `field:"optional" json:"audioSpecification" yaml:"audioSpecification"`
	// dtmf_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/lexv2models_intent#dtmf_specification Lexv2ModelsIntent#dtmf_specification}
	DtmfSpecification interface{} `field:"optional" json:"dtmfSpecification" yaml:"dtmfSpecification"`
}

