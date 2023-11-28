// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ebsvolume

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EbsVolumeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ebs_volume#availability_zone EbsVolume#availability_zone}.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ebs_volume#encrypted EbsVolume#encrypted}.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ebs_volume#final_snapshot EbsVolume#final_snapshot}.
	FinalSnapshot interface{} `field:"optional" json:"finalSnapshot" yaml:"finalSnapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ebs_volume#id EbsVolume#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ebs_volume#iops EbsVolume#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ebs_volume#kms_key_id EbsVolume#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ebs_volume#multi_attach_enabled EbsVolume#multi_attach_enabled}.
	MultiAttachEnabled interface{} `field:"optional" json:"multiAttachEnabled" yaml:"multiAttachEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ebs_volume#outpost_arn EbsVolume#outpost_arn}.
	OutpostArn *string `field:"optional" json:"outpostArn" yaml:"outpostArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ebs_volume#size EbsVolume#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ebs_volume#snapshot_id EbsVolume#snapshot_id}.
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ebs_volume#tags EbsVolume#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ebs_volume#tags_all EbsVolume#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ebs_volume#throughput EbsVolume#throughput}.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ebs_volume#timeouts EbsVolume#timeouts}
	Timeouts *EbsVolumeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/ebs_volume#type EbsVolume#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

