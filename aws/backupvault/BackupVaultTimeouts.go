// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupvault


type BackupVaultTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.44.0/docs/resources/backup_vault#delete BackupVault#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

