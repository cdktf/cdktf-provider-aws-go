// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/kinesisanalyticsv2_application#property_group_id Kinesisanalyticsv2Application#property_group_id}.
	PropertyGroupId *string `field:"required" json:"propertyGroupId" yaml:"propertyGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/kinesisanalyticsv2_application#property_map Kinesisanalyticsv2Application#property_map}.
	PropertyMap *map[string]*string `field:"required" json:"propertyMap" yaml:"propertyMap"`
}

