// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsdbsnapshot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsDbSnapshotConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/db_snapshot#db_instance_identifier DataAwsDbSnapshot#db_instance_identifier}.
	DbInstanceIdentifier *string `field:"optional" json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/db_snapshot#db_snapshot_identifier DataAwsDbSnapshot#db_snapshot_identifier}.
	DbSnapshotIdentifier *string `field:"optional" json:"dbSnapshotIdentifier" yaml:"dbSnapshotIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/db_snapshot#id DataAwsDbSnapshot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/db_snapshot#include_public DataAwsDbSnapshot#include_public}.
	IncludePublic interface{} `field:"optional" json:"includePublic" yaml:"includePublic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/db_snapshot#include_shared DataAwsDbSnapshot#include_shared}.
	IncludeShared interface{} `field:"optional" json:"includeShared" yaml:"includeShared"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/db_snapshot#most_recent DataAwsDbSnapshot#most_recent}.
	MostRecent interface{} `field:"optional" json:"mostRecent" yaml:"mostRecent"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/db_snapshot#region DataAwsDbSnapshot#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/db_snapshot#snapshot_type DataAwsDbSnapshot#snapshot_type}.
	SnapshotType *string `field:"optional" json:"snapshotType" yaml:"snapshotType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/db_snapshot#tags DataAwsDbSnapshot#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

