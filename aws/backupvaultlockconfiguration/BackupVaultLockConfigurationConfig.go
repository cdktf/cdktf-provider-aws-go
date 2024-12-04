// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupvaultlockconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupVaultLockConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/backup_vault_lock_configuration#backup_vault_name BackupVaultLockConfiguration#backup_vault_name}.
	BackupVaultName *string `field:"required" json:"backupVaultName" yaml:"backupVaultName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/backup_vault_lock_configuration#changeable_for_days BackupVaultLockConfiguration#changeable_for_days}.
	ChangeableForDays *float64 `field:"optional" json:"changeableForDays" yaml:"changeableForDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/backup_vault_lock_configuration#id BackupVaultLockConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/backup_vault_lock_configuration#max_retention_days BackupVaultLockConfiguration#max_retention_days}.
	MaxRetentionDays *float64 `field:"optional" json:"maxRetentionDays" yaml:"maxRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/backup_vault_lock_configuration#min_retention_days BackupVaultLockConfiguration#min_retention_days}.
	MinRetentionDays *float64 `field:"optional" json:"minRetentionDays" yaml:"minRetentionDays"`
}

