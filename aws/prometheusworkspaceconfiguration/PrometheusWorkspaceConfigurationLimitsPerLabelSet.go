// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package prometheusworkspaceconfiguration


type PrometheusWorkspaceConfigurationLimitsPerLabelSet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/prometheus_workspace_configuration#label_set PrometheusWorkspaceConfiguration#label_set}.
	LabelSet *map[string]*string `field:"required" json:"labelSet" yaml:"labelSet"`
	// limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/prometheus_workspace_configuration#limits PrometheusWorkspaceConfiguration#limits}
	Limits interface{} `field:"optional" json:"limits" yaml:"limits"`
}

