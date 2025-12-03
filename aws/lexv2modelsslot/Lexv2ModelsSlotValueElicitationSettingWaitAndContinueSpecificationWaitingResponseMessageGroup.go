// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsslot


type Lexv2ModelsSlotValueElicitationSettingWaitAndContinueSpecificationWaitingResponseMessageGroup struct {
	// message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lexv2models_slot#message Lexv2ModelsSlot#message}
	Message interface{} `field:"optional" json:"message" yaml:"message"`
	// variation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lexv2models_slot#variation Lexv2ModelsSlot#variation}
	Variation interface{} `field:"optional" json:"variation" yaml:"variation"`
}

