// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupBlueGreenDeploymentConfig struct {
	// deployment_ready_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/codedeploy_deployment_group#deployment_ready_option CodedeployDeploymentGroup#deployment_ready_option}
	DeploymentReadyOption *CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption `field:"optional" json:"deploymentReadyOption" yaml:"deploymentReadyOption"`
	// green_fleet_provisioning_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/codedeploy_deployment_group#green_fleet_provisioning_option CodedeployDeploymentGroup#green_fleet_provisioning_option}
	GreenFleetProvisioningOption *CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption `field:"optional" json:"greenFleetProvisioningOption" yaml:"greenFleetProvisioningOption"`
	// terminate_blue_instances_on_deployment_success block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/codedeploy_deployment_group#terminate_blue_instances_on_deployment_success CodedeployDeploymentGroup#terminate_blue_instances_on_deployment_success}
	TerminateBlueInstancesOnDeploymentSuccess *CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess `field:"optional" json:"terminateBlueInstancesOnDeploymentSuccess" yaml:"terminateBlueInstancesOnDeploymentSuccess"`
}

