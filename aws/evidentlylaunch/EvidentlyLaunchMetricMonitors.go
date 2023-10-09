// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package evidentlylaunch


type EvidentlyLaunchMetricMonitors struct {
	// metric_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/evidently_launch#metric_definition EvidentlyLaunch#metric_definition}
	MetricDefinition *EvidentlyLaunchMetricMonitorsMetricDefinition `field:"required" json:"metricDefinition" yaml:"metricDefinition"`
}

