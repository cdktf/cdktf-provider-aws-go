// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsslot


type Lexv2ModelsSlotValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponseMessageGroupMessage struct {
	// custom_playload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lexv2models_slot#custom_playload Lexv2ModelsSlot#custom_playload}
	CustomPlayload interface{} `field:"optional" json:"customPlayload" yaml:"customPlayload"`
	// image_response_card block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lexv2models_slot#image_response_card Lexv2ModelsSlot#image_response_card}
	ImageResponseCard interface{} `field:"optional" json:"imageResponseCard" yaml:"imageResponseCard"`
	// plain_text_message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lexv2models_slot#plain_text_message Lexv2ModelsSlot#plain_text_message}
	PlainTextMessage interface{} `field:"optional" json:"plainTextMessage" yaml:"plainTextMessage"`
	// ssml_message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lexv2models_slot#ssml_message Lexv2ModelsSlot#ssml_message}
	SsmlMessage interface{} `field:"optional" json:"ssmlMessage" yaml:"ssmlMessage"`
}

