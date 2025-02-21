// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory


type WorkspacesDirectoryWorkspaceAccessProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/workspaces_directory#device_type_android WorkspacesDirectory#device_type_android}.
	DeviceTypeAndroid *string `field:"optional" json:"deviceTypeAndroid" yaml:"deviceTypeAndroid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/workspaces_directory#device_type_chromeos WorkspacesDirectory#device_type_chromeos}.
	DeviceTypeChromeos *string `field:"optional" json:"deviceTypeChromeos" yaml:"deviceTypeChromeos"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/workspaces_directory#device_type_ios WorkspacesDirectory#device_type_ios}.
	DeviceTypeIos *string `field:"optional" json:"deviceTypeIos" yaml:"deviceTypeIos"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/workspaces_directory#device_type_linux WorkspacesDirectory#device_type_linux}.
	DeviceTypeLinux *string `field:"optional" json:"deviceTypeLinux" yaml:"deviceTypeLinux"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/workspaces_directory#device_type_osx WorkspacesDirectory#device_type_osx}.
	DeviceTypeOsx *string `field:"optional" json:"deviceTypeOsx" yaml:"deviceTypeOsx"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/workspaces_directory#device_type_web WorkspacesDirectory#device_type_web}.
	DeviceTypeWeb *string `field:"optional" json:"deviceTypeWeb" yaml:"deviceTypeWeb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/workspaces_directory#device_type_windows WorkspacesDirectory#device_type_windows}.
	DeviceTypeWindows *string `field:"optional" json:"deviceTypeWindows" yaml:"deviceTypeWindows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/workspaces_directory#device_type_zeroclient WorkspacesDirectory#device_type_zeroclient}.
	DeviceTypeZeroclient *string `field:"optional" json:"deviceTypeZeroclient" yaml:"deviceTypeZeroclient"`
}

