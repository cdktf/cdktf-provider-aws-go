// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentFulfillmentCodeHook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/lexv2models_intent#enabled Lexv2ModelsIntent#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/lexv2models_intent#active Lexv2ModelsIntent#active}.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// fulfillment_updates_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/lexv2models_intent#fulfillment_updates_specification Lexv2ModelsIntent#fulfillment_updates_specification}
	FulfillmentUpdatesSpecification interface{} `field:"optional" json:"fulfillmentUpdatesSpecification" yaml:"fulfillmentUpdatesSpecification"`
	// post_fulfillment_status_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/lexv2models_intent#post_fulfillment_status_specification Lexv2ModelsIntent#post_fulfillment_status_specification}
	PostFulfillmentStatusSpecification interface{} `field:"optional" json:"postFulfillmentStatusSpecification" yaml:"postFulfillmentStatusSpecification"`
}

