// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecscluster


type EcsClusterConfiguration struct {
	// execute_command_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/ecs_cluster#execute_command_configuration EcsCluster#execute_command_configuration}
	ExecuteCommandConfiguration *EcsClusterConfigurationExecuteCommandConfiguration `field:"optional" json:"executeCommandConfiguration" yaml:"executeCommandConfiguration"`
	// managed_storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/ecs_cluster#managed_storage_configuration EcsCluster#managed_storage_configuration}
	ManagedStorageConfiguration *EcsClusterConfigurationManagedStorageConfiguration `field:"optional" json:"managedStorageConfiguration" yaml:"managedStorageConfiguration"`
}

