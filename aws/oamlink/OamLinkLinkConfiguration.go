// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oamlink


type OamLinkLinkConfiguration struct {
	// log_group_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/oam_link#log_group_configuration OamLink#log_group_configuration}
	LogGroupConfiguration *OamLinkLinkConfigurationLogGroupConfiguration `field:"optional" json:"logGroupConfiguration" yaml:"logGroupConfiguration"`
	// metric_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/oam_link#metric_configuration OamLink#metric_configuration}
	MetricConfiguration *OamLinkLinkConfigurationMetricConfiguration `field:"optional" json:"metricConfiguration" yaml:"metricConfiguration"`
}

