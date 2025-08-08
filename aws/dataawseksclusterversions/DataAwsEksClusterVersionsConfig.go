// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawseksclusterversions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsEksClusterVersionsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/eks_cluster_versions#cluster_type DataAwsEksClusterVersions#cluster_type}.
	ClusterType *string `field:"optional" json:"clusterType" yaml:"clusterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/eks_cluster_versions#cluster_versions_only DataAwsEksClusterVersions#cluster_versions_only}.
	ClusterVersionsOnly *[]*string `field:"optional" json:"clusterVersionsOnly" yaml:"clusterVersionsOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/eks_cluster_versions#default_only DataAwsEksClusterVersions#default_only}.
	DefaultOnly interface{} `field:"optional" json:"defaultOnly" yaml:"defaultOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/eks_cluster_versions#include_all DataAwsEksClusterVersions#include_all}.
	IncludeAll interface{} `field:"optional" json:"includeAll" yaml:"includeAll"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/eks_cluster_versions#region DataAwsEksClusterVersions#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/eks_cluster_versions#version_status DataAwsEksClusterVersions#version_status}.
	VersionStatus *string `field:"optional" json:"versionStatus" yaml:"versionStatus"`
}

