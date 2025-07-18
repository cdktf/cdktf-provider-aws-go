// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsslot


type Lexv2ModelsSlotValueElicitationSettingWaitAndContinueSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/lexv2models_slot#active Lexv2ModelsSlot#active}.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// continue_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/lexv2models_slot#continue_response Lexv2ModelsSlot#continue_response}
	ContinueResponse interface{} `field:"optional" json:"continueResponse" yaml:"continueResponse"`
	// still_waiting_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/lexv2models_slot#still_waiting_response Lexv2ModelsSlot#still_waiting_response}
	StillWaitingResponse interface{} `field:"optional" json:"stillWaitingResponse" yaml:"stillWaitingResponse"`
	// waiting_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/lexv2models_slot#waiting_response Lexv2ModelsSlot#waiting_response}
	WaitingResponse interface{} `field:"optional" json:"waitingResponse" yaml:"waitingResponse"`
}

