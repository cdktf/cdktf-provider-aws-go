// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmmaintenancewindowtask


type SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/ssm_maintenance_window_task#cloudwatch_log_group_name SsmMaintenanceWindowTask#cloudwatch_log_group_name}.
	CloudwatchLogGroupName *string `field:"optional" json:"cloudwatchLogGroupName" yaml:"cloudwatchLogGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/ssm_maintenance_window_task#cloudwatch_output_enabled SsmMaintenanceWindowTask#cloudwatch_output_enabled}.
	CloudwatchOutputEnabled interface{} `field:"optional" json:"cloudwatchOutputEnabled" yaml:"cloudwatchOutputEnabled"`
}

