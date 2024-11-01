// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package evidentlylaunch


type EvidentlyLaunchMetricMonitorsMetricDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/evidently_launch#entity_id_key EvidentlyLaunch#entity_id_key}.
	EntityIdKey *string `field:"required" json:"entityIdKey" yaml:"entityIdKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/evidently_launch#name EvidentlyLaunch#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/evidently_launch#value_key EvidentlyLaunch#value_key}.
	ValueKey *string `field:"required" json:"valueKey" yaml:"valueKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/evidently_launch#event_pattern EvidentlyLaunch#event_pattern}.
	EventPattern *string `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/evidently_launch#unit_label EvidentlyLaunch#unit_label}.
	UnitLabel *string `field:"optional" json:"unitLabel" yaml:"unitLabel"`
}

