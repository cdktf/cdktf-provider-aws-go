// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsScheduleArchiveRuleArchiveRetainRule struct {
	// retention_archive_tier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/dlm_lifecycle_policy#retention_archive_tier DlmLifecyclePolicy#retention_archive_tier}
	RetentionArchiveTier *DlmLifecyclePolicyPolicyDetailsScheduleArchiveRuleArchiveRetainRuleRetentionArchiveTier `field:"required" json:"retentionArchiveTier" yaml:"retentionArchiveTier"`
}

