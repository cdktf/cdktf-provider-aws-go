// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rbinrule


type RbinRuleLockConfiguration struct {
	// unlock_delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/rbin_rule#unlock_delay RbinRule#unlock_delay}
	UnlockDelay *RbinRuleLockConfigurationUnlockDelay `field:"required" json:"unlockDelay" yaml:"unlockDelay"`
}

