// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspacesworkspace


type WorkspacesWorkspaceWorkspaceProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/workspaces_workspace#compute_type_name WorkspacesWorkspace#compute_type_name}.
	ComputeTypeName *string `field:"optional" json:"computeTypeName" yaml:"computeTypeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/workspaces_workspace#root_volume_size_gib WorkspacesWorkspace#root_volume_size_gib}.
	RootVolumeSizeGib *float64 `field:"optional" json:"rootVolumeSizeGib" yaml:"rootVolumeSizeGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/workspaces_workspace#running_mode WorkspacesWorkspace#running_mode}.
	RunningMode *string `field:"optional" json:"runningMode" yaml:"runningMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/workspaces_workspace#running_mode_auto_stop_timeout_in_minutes WorkspacesWorkspace#running_mode_auto_stop_timeout_in_minutes}.
	RunningModeAutoStopTimeoutInMinutes *float64 `field:"optional" json:"runningModeAutoStopTimeoutInMinutes" yaml:"runningModeAutoStopTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/workspaces_workspace#user_volume_size_gib WorkspacesWorkspace#user_volume_size_gib}.
	UserVolumeSizeGib *float64 `field:"optional" json:"userVolumeSizeGib" yaml:"userVolumeSizeGib"`
}

