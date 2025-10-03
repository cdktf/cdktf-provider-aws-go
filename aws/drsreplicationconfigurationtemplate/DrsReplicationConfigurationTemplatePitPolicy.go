// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package drsreplicationconfigurationtemplate


type DrsReplicationConfigurationTemplatePitPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/drs_replication_configuration_template#interval DrsReplicationConfigurationTemplate#interval}.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/drs_replication_configuration_template#retention_duration DrsReplicationConfigurationTemplate#retention_duration}.
	RetentionDuration *float64 `field:"required" json:"retentionDuration" yaml:"retentionDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/drs_replication_configuration_template#units DrsReplicationConfigurationTemplate#units}.
	Units *string `field:"required" json:"units" yaml:"units"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/drs_replication_configuration_template#enabled DrsReplicationConfigurationTemplate#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/drs_replication_configuration_template#rule_id DrsReplicationConfigurationTemplate#rule_id}.
	RuleId *float64 `field:"optional" json:"ruleId" yaml:"ruleId"`
}

