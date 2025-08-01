// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmmaintenancewindowtask


type SsmMaintenanceWindowTaskTaskInvocationParameters struct {
	// automation_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ssm_maintenance_window_task#automation_parameters SsmMaintenanceWindowTask#automation_parameters}
	AutomationParameters *SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters `field:"optional" json:"automationParameters" yaml:"automationParameters"`
	// lambda_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ssm_maintenance_window_task#lambda_parameters SsmMaintenanceWindowTask#lambda_parameters}
	LambdaParameters *SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters `field:"optional" json:"lambdaParameters" yaml:"lambdaParameters"`
	// run_command_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ssm_maintenance_window_task#run_command_parameters SsmMaintenanceWindowTask#run_command_parameters}
	RunCommandParameters *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters `field:"optional" json:"runCommandParameters" yaml:"runCommandParameters"`
	// step_functions_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ssm_maintenance_window_task#step_functions_parameters SsmMaintenanceWindowTask#step_functions_parameters}
	StepFunctionsParameters *SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters `field:"optional" json:"stepFunctionsParameters" yaml:"stepFunctionsParameters"`
}

