// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxwindowsfilesystem


type FsxWindowsFileSystemAuditLogConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/fsx_windows_file_system#audit_log_destination FsxWindowsFileSystem#audit_log_destination}.
	AuditLogDestination *string `field:"optional" json:"auditLogDestination" yaml:"auditLogDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/fsx_windows_file_system#file_access_audit_log_level FsxWindowsFileSystem#file_access_audit_log_level}.
	FileAccessAuditLogLevel *string `field:"optional" json:"fileAccessAuditLogLevel" yaml:"fileAccessAuditLogLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/fsx_windows_file_system#file_share_access_audit_log_level FsxWindowsFileSystem#file_share_access_audit_log_level}.
	FileShareAccessAuditLogLevel *string `field:"optional" json:"fileShareAccessAuditLogLevel" yaml:"fileShareAccessAuditLogLevel"`
}

