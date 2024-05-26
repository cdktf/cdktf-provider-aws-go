// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotindexingconfiguration


type IotIndexingConfigurationThingGroupIndexingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/iot_indexing_configuration#thing_group_indexing_mode IotIndexingConfiguration#thing_group_indexing_mode}.
	ThingGroupIndexingMode *string `field:"required" json:"thingGroupIndexingMode" yaml:"thingGroupIndexingMode"`
	// custom_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/iot_indexing_configuration#custom_field IotIndexingConfiguration#custom_field}
	CustomField interface{} `field:"optional" json:"customField" yaml:"customField"`
	// managed_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/iot_indexing_configuration#managed_field IotIndexingConfiguration#managed_field}
	ManagedField interface{} `field:"optional" json:"managedField" yaml:"managedField"`
}

