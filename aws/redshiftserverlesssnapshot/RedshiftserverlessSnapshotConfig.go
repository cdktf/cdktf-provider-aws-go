// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshiftserverlesssnapshot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedshiftserverlessSnapshotConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/redshiftserverless_snapshot#namespace_name RedshiftserverlessSnapshot#namespace_name}.
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/redshiftserverless_snapshot#snapshot_name RedshiftserverlessSnapshot#snapshot_name}.
	SnapshotName *string `field:"required" json:"snapshotName" yaml:"snapshotName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/redshiftserverless_snapshot#id RedshiftserverlessSnapshot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/redshiftserverless_snapshot#retention_period RedshiftserverlessSnapshot#retention_period}.
	RetentionPeriod *float64 `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
}

