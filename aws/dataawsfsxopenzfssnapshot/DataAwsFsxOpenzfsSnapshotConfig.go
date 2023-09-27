// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsfsxopenzfssnapshot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsFsxOpenzfsSnapshotConfig struct {
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
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/data-sources/fsx_openzfs_snapshot#filter DataAwsFsxOpenzfsSnapshot#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/data-sources/fsx_openzfs_snapshot#id DataAwsFsxOpenzfsSnapshot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/data-sources/fsx_openzfs_snapshot#most_recent DataAwsFsxOpenzfsSnapshot#most_recent}.
	MostRecent interface{} `field:"optional" json:"mostRecent" yaml:"mostRecent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/data-sources/fsx_openzfs_snapshot#name DataAwsFsxOpenzfsSnapshot#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/data-sources/fsx_openzfs_snapshot#snapshot_ids DataAwsFsxOpenzfsSnapshot#snapshot_ids}.
	SnapshotIds *[]*string `field:"optional" json:"snapshotIds" yaml:"snapshotIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/data-sources/fsx_openzfs_snapshot#tags DataAwsFsxOpenzfsSnapshot#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

