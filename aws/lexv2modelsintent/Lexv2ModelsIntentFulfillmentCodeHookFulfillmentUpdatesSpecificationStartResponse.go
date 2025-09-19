// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponse struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/lexv2models_intent#allow_interrupt Lexv2ModelsIntent#allow_interrupt}.
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/lexv2models_intent#delay_in_seconds Lexv2ModelsIntent#delay_in_seconds}.
	DelayInSeconds *float64 `field:"optional" json:"delayInSeconds" yaml:"delayInSeconds"`
	// message_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/lexv2models_intent#message_group Lexv2ModelsIntent#message_group}
	MessageGroup interface{} `field:"optional" json:"messageGroup" yaml:"messageGroup"`
}

