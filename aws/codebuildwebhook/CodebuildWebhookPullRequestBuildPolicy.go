// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codebuildwebhook


type CodebuildWebhookPullRequestBuildPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/codebuild_webhook#requires_comment_approval CodebuildWebhook#requires_comment_approval}.
	RequiresCommentApproval *string `field:"required" json:"requiresCommentApproval" yaml:"requiresCommentApproval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/codebuild_webhook#approver_roles CodebuildWebhook#approver_roles}.
	ApproverRoles *[]*string `field:"optional" json:"approverRoles" yaml:"approverRoles"`
}

