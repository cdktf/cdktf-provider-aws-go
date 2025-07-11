// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssoadminmanagedpolicyattachment


type SsoadminManagedPolicyAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/ssoadmin_managed_policy_attachment#create SsoadminManagedPolicyAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/ssoadmin_managed_policy_attachment#delete SsoadminManagedPolicyAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

