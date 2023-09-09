// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecscluster


type EcsClusterConfiguration struct {
	// execute_command_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/ecs_cluster#execute_command_configuration EcsCluster#execute_command_configuration}
	ExecuteCommandConfiguration *EcsClusterConfigurationExecuteCommandConfiguration `field:"optional" json:"executeCommandConfiguration" yaml:"executeCommandConfiguration"`
}

