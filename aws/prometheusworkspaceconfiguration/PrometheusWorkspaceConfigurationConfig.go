// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package prometheusworkspaceconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrometheusWorkspaceConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/prometheus_workspace_configuration#workspace_id PrometheusWorkspaceConfiguration#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// limits_per_label_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/prometheus_workspace_configuration#limits_per_label_set PrometheusWorkspaceConfiguration#limits_per_label_set}
	LimitsPerLabelSet interface{} `field:"optional" json:"limitsPerLabelSet" yaml:"limitsPerLabelSet"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/prometheus_workspace_configuration#region PrometheusWorkspaceConfiguration#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/prometheus_workspace_configuration#retention_period_in_days PrometheusWorkspaceConfiguration#retention_period_in_days}.
	RetentionPeriodInDays *float64 `field:"optional" json:"retentionPeriodInDays" yaml:"retentionPeriodInDays"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/prometheus_workspace_configuration#timeouts PrometheusWorkspaceConfiguration#timeouts}
	Timeouts *PrometheusWorkspaceConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

