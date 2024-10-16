// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package efsbackuppolicy


type EfsBackupPolicyBackupPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.0/docs/resources/efs_backup_policy#status EfsBackupPolicy#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
}

