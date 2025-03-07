// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentClosingSettingClosingResponseMessageGroupMessage struct {
	// custom_payload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/lexv2models_intent#custom_payload Lexv2ModelsIntent#custom_payload}
	CustomPayload interface{} `field:"optional" json:"customPayload" yaml:"customPayload"`
	// image_response_card block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/lexv2models_intent#image_response_card Lexv2ModelsIntent#image_response_card}
	ImageResponseCard interface{} `field:"optional" json:"imageResponseCard" yaml:"imageResponseCard"`
	// plain_text_message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/lexv2models_intent#plain_text_message Lexv2ModelsIntent#plain_text_message}
	PlainTextMessage interface{} `field:"optional" json:"plainTextMessage" yaml:"plainTextMessage"`
	// ssml_message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/lexv2models_intent#ssml_message Lexv2ModelsIntent#ssml_message}
	SsmlMessage interface{} `field:"optional" json:"ssmlMessage" yaml:"ssmlMessage"`
}

