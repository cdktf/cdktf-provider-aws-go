// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qbusinessapplication


type QbusinessApplicationAttachmentsConfiguration struct {
	// Status information about whether file upload functionality is activated or deactivated for your end user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/qbusiness_application#attachments_control_mode QbusinessApplication#attachments_control_mode}
	AttachmentsControlMode *string `field:"required" json:"attachmentsControlMode" yaml:"attachmentsControlMode"`
}

