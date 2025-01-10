// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotthinggroup


type IotThingGroupProperties struct {
	// attribute_payload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/iot_thing_group#attribute_payload IotThingGroup#attribute_payload}
	AttributePayload *IotThingGroupPropertiesAttributePayload `field:"optional" json:"attributePayload" yaml:"attributePayload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/iot_thing_group#description IotThingGroup#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

