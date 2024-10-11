// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ebssnapshotimport

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EbsSnapshotImportConfig struct {
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
	// disk_container block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/ebs_snapshot_import#disk_container EbsSnapshotImport#disk_container}
	DiskContainer *EbsSnapshotImportDiskContainer `field:"required" json:"diskContainer" yaml:"diskContainer"`
	// client_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/ebs_snapshot_import#client_data EbsSnapshotImport#client_data}
	ClientData *EbsSnapshotImportClientData `field:"optional" json:"clientData" yaml:"clientData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/ebs_snapshot_import#description EbsSnapshotImport#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/ebs_snapshot_import#encrypted EbsSnapshotImport#encrypted}.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/ebs_snapshot_import#id EbsSnapshotImport#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/ebs_snapshot_import#kms_key_id EbsSnapshotImport#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/ebs_snapshot_import#permanent_restore EbsSnapshotImport#permanent_restore}.
	PermanentRestore interface{} `field:"optional" json:"permanentRestore" yaml:"permanentRestore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/ebs_snapshot_import#role_name EbsSnapshotImport#role_name}.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/ebs_snapshot_import#storage_tier EbsSnapshotImport#storage_tier}.
	StorageTier *string `field:"optional" json:"storageTier" yaml:"storageTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/ebs_snapshot_import#tags EbsSnapshotImport#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/ebs_snapshot_import#tags_all EbsSnapshotImport#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/ebs_snapshot_import#temporary_restore_days EbsSnapshotImport#temporary_restore_days}.
	TemporaryRestoreDays *float64 `field:"optional" json:"temporaryRestoreDays" yaml:"temporaryRestoreDays"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/ebs_snapshot_import#timeouts EbsSnapshotImport#timeouts}
	Timeouts *EbsSnapshotImportTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

