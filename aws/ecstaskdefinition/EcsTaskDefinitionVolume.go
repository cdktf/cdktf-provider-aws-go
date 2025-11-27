// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionVolume struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_task_definition#name EcsTaskDefinition#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_task_definition#configure_at_launch EcsTaskDefinition#configure_at_launch}.
	ConfigureAtLaunch interface{} `field:"optional" json:"configureAtLaunch" yaml:"configureAtLaunch"`
	// docker_volume_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_task_definition#docker_volume_configuration EcsTaskDefinition#docker_volume_configuration}
	DockerVolumeConfiguration *EcsTaskDefinitionVolumeDockerVolumeConfiguration `field:"optional" json:"dockerVolumeConfiguration" yaml:"dockerVolumeConfiguration"`
	// efs_volume_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_task_definition#efs_volume_configuration EcsTaskDefinition#efs_volume_configuration}
	EfsVolumeConfiguration *EcsTaskDefinitionVolumeEfsVolumeConfiguration `field:"optional" json:"efsVolumeConfiguration" yaml:"efsVolumeConfiguration"`
	// fsx_windows_file_server_volume_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_task_definition#fsx_windows_file_server_volume_configuration EcsTaskDefinition#fsx_windows_file_server_volume_configuration}
	FsxWindowsFileServerVolumeConfiguration *EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfiguration `field:"optional" json:"fsxWindowsFileServerVolumeConfiguration" yaml:"fsxWindowsFileServerVolumeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_task_definition#host_path EcsTaskDefinition#host_path}.
	HostPath *string `field:"optional" json:"hostPath" yaml:"hostPath"`
}

