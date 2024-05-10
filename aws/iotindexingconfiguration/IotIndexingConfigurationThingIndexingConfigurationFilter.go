// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotindexingconfiguration


type IotIndexingConfigurationThingIndexingConfigurationFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/iot_indexing_configuration#named_shadow_names IotIndexingConfiguration#named_shadow_names}.
	NamedShadowNames *[]*string `field:"optional" json:"namedShadowNames" yaml:"namedShadowNames"`
}

