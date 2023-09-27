// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory


type WorkspacesDirectorySelfServicePermissions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/workspaces_directory#change_compute_type WorkspacesDirectory#change_compute_type}.
	ChangeComputeType interface{} `field:"optional" json:"changeComputeType" yaml:"changeComputeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/workspaces_directory#increase_volume_size WorkspacesDirectory#increase_volume_size}.
	IncreaseVolumeSize interface{} `field:"optional" json:"increaseVolumeSize" yaml:"increaseVolumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/workspaces_directory#rebuild_workspace WorkspacesDirectory#rebuild_workspace}.
	RebuildWorkspace interface{} `field:"optional" json:"rebuildWorkspace" yaml:"rebuildWorkspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/workspaces_directory#restart_workspace WorkspacesDirectory#restart_workspace}.
	RestartWorkspace interface{} `field:"optional" json:"restartWorkspace" yaml:"restartWorkspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/workspaces_directory#switch_running_mode WorkspacesDirectory#switch_running_mode}.
	SwitchRunningMode interface{} `field:"optional" json:"switchRunningMode" yaml:"switchRunningMode"`
}

