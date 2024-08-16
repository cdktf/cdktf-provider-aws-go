// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/codedeploy_deployment_group#action CodedeployDeploymentGroup#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/codedeploy_deployment_group#termination_wait_time_in_minutes CodedeployDeploymentGroup#termination_wait_time_in_minutes}.
	TerminationWaitTimeInMinutes *float64 `field:"optional" json:"terminationWaitTimeInMinutes" yaml:"terminationWaitTimeInMinutes"`
}

