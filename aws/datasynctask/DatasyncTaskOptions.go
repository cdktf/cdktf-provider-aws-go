// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasynctask


type DatasyncTaskOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/datasync_task#atime DatasyncTask#atime}.
	Atime *string `field:"optional" json:"atime" yaml:"atime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/datasync_task#bytes_per_second DatasyncTask#bytes_per_second}.
	BytesPerSecond *float64 `field:"optional" json:"bytesPerSecond" yaml:"bytesPerSecond"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/datasync_task#gid DatasyncTask#gid}.
	Gid *string `field:"optional" json:"gid" yaml:"gid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/datasync_task#log_level DatasyncTask#log_level}.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/datasync_task#mtime DatasyncTask#mtime}.
	Mtime *string `field:"optional" json:"mtime" yaml:"mtime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/datasync_task#object_tags DatasyncTask#object_tags}.
	ObjectTags *string `field:"optional" json:"objectTags" yaml:"objectTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/datasync_task#overwrite_mode DatasyncTask#overwrite_mode}.
	OverwriteMode *string `field:"optional" json:"overwriteMode" yaml:"overwriteMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/datasync_task#posix_permissions DatasyncTask#posix_permissions}.
	PosixPermissions *string `field:"optional" json:"posixPermissions" yaml:"posixPermissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/datasync_task#preserve_deleted_files DatasyncTask#preserve_deleted_files}.
	PreserveDeletedFiles *string `field:"optional" json:"preserveDeletedFiles" yaml:"preserveDeletedFiles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/datasync_task#preserve_devices DatasyncTask#preserve_devices}.
	PreserveDevices *string `field:"optional" json:"preserveDevices" yaml:"preserveDevices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/datasync_task#security_descriptor_copy_flags DatasyncTask#security_descriptor_copy_flags}.
	SecurityDescriptorCopyFlags *string `field:"optional" json:"securityDescriptorCopyFlags" yaml:"securityDescriptorCopyFlags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/datasync_task#task_queueing DatasyncTask#task_queueing}.
	TaskQueueing *string `field:"optional" json:"taskQueueing" yaml:"taskQueueing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/datasync_task#transfer_mode DatasyncTask#transfer_mode}.
	TransferMode *string `field:"optional" json:"transferMode" yaml:"transferMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/datasync_task#uid DatasyncTask#uid}.
	Uid *string `field:"optional" json:"uid" yaml:"uid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/datasync_task#verify_mode DatasyncTask#verify_mode}.
	VerifyMode *string `field:"optional" json:"verifyMode" yaml:"verifyMode"`
}

