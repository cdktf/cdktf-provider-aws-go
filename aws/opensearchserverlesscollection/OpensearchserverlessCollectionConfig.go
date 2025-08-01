// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchserverlesscollection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpensearchserverlessCollectionConfig struct {
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
	// Name of the collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/opensearchserverless_collection#name OpensearchserverlessCollection#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Description of the collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/opensearchserverless_collection#description OpensearchserverlessCollection#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/opensearchserverless_collection#region OpensearchserverlessCollection#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Indicates whether standby replicas should be used for a collection. One of `ENABLED` or `DISABLED`. Defaults to `ENABLED`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/opensearchserverless_collection#standby_replicas OpensearchserverlessCollection#standby_replicas}
	StandbyReplicas *string `field:"optional" json:"standbyReplicas" yaml:"standbyReplicas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/opensearchserverless_collection#tags OpensearchserverlessCollection#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/opensearchserverless_collection#timeouts OpensearchserverlessCollection#timeouts}
	Timeouts *OpensearchserverlessCollectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Type of collection. One of `SEARCH`, `TIMESERIES`, or `VECTORSEARCH`. Defaults to `TIMESERIES`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/opensearchserverless_collection#type OpensearchserverlessCollection#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

