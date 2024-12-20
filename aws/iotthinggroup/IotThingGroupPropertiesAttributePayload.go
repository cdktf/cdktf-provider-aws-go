// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotthinggroup


type IotThingGroupPropertiesAttributePayload struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/iot_thing_group#attributes IotThingGroup#attributes}.
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
}

