// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotindexingconfiguration


type IotIndexingConfigurationThingIndexingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/iot_indexing_configuration#thing_indexing_mode IotIndexingConfiguration#thing_indexing_mode}.
	ThingIndexingMode *string `field:"required" json:"thingIndexingMode" yaml:"thingIndexingMode"`
	// custom_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/iot_indexing_configuration#custom_field IotIndexingConfiguration#custom_field}
	CustomField interface{} `field:"optional" json:"customField" yaml:"customField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/iot_indexing_configuration#device_defender_indexing_mode IotIndexingConfiguration#device_defender_indexing_mode}.
	DeviceDefenderIndexingMode *string `field:"optional" json:"deviceDefenderIndexingMode" yaml:"deviceDefenderIndexingMode"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/iot_indexing_configuration#filter IotIndexingConfiguration#filter}
	Filter *IotIndexingConfigurationThingIndexingConfigurationFilter `field:"optional" json:"filter" yaml:"filter"`
	// managed_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/iot_indexing_configuration#managed_field IotIndexingConfiguration#managed_field}
	ManagedField interface{} `field:"optional" json:"managedField" yaml:"managedField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/iot_indexing_configuration#named_shadow_indexing_mode IotIndexingConfiguration#named_shadow_indexing_mode}.
	NamedShadowIndexingMode *string `field:"optional" json:"namedShadowIndexingMode" yaml:"namedShadowIndexingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/iot_indexing_configuration#thing_connectivity_indexing_mode IotIndexingConfiguration#thing_connectivity_indexing_mode}.
	ThingConnectivityIndexingMode *string `field:"optional" json:"thingConnectivityIndexingMode" yaml:"thingConnectivityIndexingMode"`
}

