// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxopenzfsvolume

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FsxOpenzfsVolumeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/fsx_openzfs_volume#name FsxOpenzfsVolume#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/fsx_openzfs_volume#parent_volume_id FsxOpenzfsVolume#parent_volume_id}.
	ParentVolumeId *string `field:"required" json:"parentVolumeId" yaml:"parentVolumeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/fsx_openzfs_volume#copy_tags_to_snapshots FsxOpenzfsVolume#copy_tags_to_snapshots}.
	CopyTagsToSnapshots interface{} `field:"optional" json:"copyTagsToSnapshots" yaml:"copyTagsToSnapshots"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/fsx_openzfs_volume#data_compression_type FsxOpenzfsVolume#data_compression_type}.
	DataCompressionType *string `field:"optional" json:"dataCompressionType" yaml:"dataCompressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/fsx_openzfs_volume#delete_volume_options FsxOpenzfsVolume#delete_volume_options}.
	DeleteVolumeOptions *[]*string `field:"optional" json:"deleteVolumeOptions" yaml:"deleteVolumeOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/fsx_openzfs_volume#id FsxOpenzfsVolume#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// nfs_exports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/fsx_openzfs_volume#nfs_exports FsxOpenzfsVolume#nfs_exports}
	NfsExports *FsxOpenzfsVolumeNfsExports `field:"optional" json:"nfsExports" yaml:"nfsExports"`
	// origin_snapshot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/fsx_openzfs_volume#origin_snapshot FsxOpenzfsVolume#origin_snapshot}
	OriginSnapshot *FsxOpenzfsVolumeOriginSnapshot `field:"optional" json:"originSnapshot" yaml:"originSnapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/fsx_openzfs_volume#read_only FsxOpenzfsVolume#read_only}.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/fsx_openzfs_volume#record_size_kib FsxOpenzfsVolume#record_size_kib}.
	RecordSizeKib *float64 `field:"optional" json:"recordSizeKib" yaml:"recordSizeKib"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/fsx_openzfs_volume#region FsxOpenzfsVolume#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/fsx_openzfs_volume#storage_capacity_quota_gib FsxOpenzfsVolume#storage_capacity_quota_gib}.
	StorageCapacityQuotaGib *float64 `field:"optional" json:"storageCapacityQuotaGib" yaml:"storageCapacityQuotaGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/fsx_openzfs_volume#storage_capacity_reservation_gib FsxOpenzfsVolume#storage_capacity_reservation_gib}.
	StorageCapacityReservationGib *float64 `field:"optional" json:"storageCapacityReservationGib" yaml:"storageCapacityReservationGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/fsx_openzfs_volume#tags FsxOpenzfsVolume#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/fsx_openzfs_volume#tags_all FsxOpenzfsVolume#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/fsx_openzfs_volume#timeouts FsxOpenzfsVolume#timeouts}
	Timeouts *FsxOpenzfsVolumeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// user_and_group_quotas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/fsx_openzfs_volume#user_and_group_quotas FsxOpenzfsVolume#user_and_group_quotas}
	UserAndGroupQuotas interface{} `field:"optional" json:"userAndGroupQuotas" yaml:"userAndGroupQuotas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/fsx_openzfs_volume#volume_type FsxOpenzfsVolume#volume_type}.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

