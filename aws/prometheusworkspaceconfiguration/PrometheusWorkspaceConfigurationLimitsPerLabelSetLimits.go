// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package prometheusworkspaceconfiguration


type PrometheusWorkspaceConfigurationLimitsPerLabelSetLimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/prometheus_workspace_configuration#max_series PrometheusWorkspaceConfiguration#max_series}.
	MaxSeries *float64 `field:"required" json:"maxSeries" yaml:"maxSeries"`
}

