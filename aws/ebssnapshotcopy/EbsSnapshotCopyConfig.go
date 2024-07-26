// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ebssnapshotcopy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EbsSnapshotCopyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/ebs_snapshot_copy#source_region EbsSnapshotCopy#source_region}.
	SourceRegion *string `field:"required" json:"sourceRegion" yaml:"sourceRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/ebs_snapshot_copy#source_snapshot_id EbsSnapshotCopy#source_snapshot_id}.
	SourceSnapshotId *string `field:"required" json:"sourceSnapshotId" yaml:"sourceSnapshotId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/ebs_snapshot_copy#description EbsSnapshotCopy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/ebs_snapshot_copy#encrypted EbsSnapshotCopy#encrypted}.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/ebs_snapshot_copy#id EbsSnapshotCopy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/ebs_snapshot_copy#kms_key_id EbsSnapshotCopy#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/ebs_snapshot_copy#permanent_restore EbsSnapshotCopy#permanent_restore}.
	PermanentRestore interface{} `field:"optional" json:"permanentRestore" yaml:"permanentRestore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/ebs_snapshot_copy#storage_tier EbsSnapshotCopy#storage_tier}.
	StorageTier *string `field:"optional" json:"storageTier" yaml:"storageTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/ebs_snapshot_copy#tags EbsSnapshotCopy#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/ebs_snapshot_copy#tags_all EbsSnapshotCopy#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/ebs_snapshot_copy#temporary_restore_days EbsSnapshotCopy#temporary_restore_days}.
	TemporaryRestoreDays *float64 `field:"optional" json:"temporaryRestoreDays" yaml:"temporaryRestoreDays"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/ebs_snapshot_copy#timeouts EbsSnapshotCopy#timeouts}
	Timeouts *EbsSnapshotCopyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

