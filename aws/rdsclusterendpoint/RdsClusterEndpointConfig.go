// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdsclusterendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsClusterEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/rds_cluster_endpoint#cluster_endpoint_identifier RdsClusterEndpoint#cluster_endpoint_identifier}.
	ClusterEndpointIdentifier *string `field:"required" json:"clusterEndpointIdentifier" yaml:"clusterEndpointIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/rds_cluster_endpoint#cluster_identifier RdsClusterEndpoint#cluster_identifier}.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/rds_cluster_endpoint#custom_endpoint_type RdsClusterEndpoint#custom_endpoint_type}.
	CustomEndpointType *string `field:"required" json:"customEndpointType" yaml:"customEndpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/rds_cluster_endpoint#excluded_members RdsClusterEndpoint#excluded_members}.
	ExcludedMembers *[]*string `field:"optional" json:"excludedMembers" yaml:"excludedMembers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/rds_cluster_endpoint#id RdsClusterEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/rds_cluster_endpoint#region RdsClusterEndpoint#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/rds_cluster_endpoint#static_members RdsClusterEndpoint#static_members}.
	StaticMembers *[]*string `field:"optional" json:"staticMembers" yaml:"staticMembers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/rds_cluster_endpoint#tags RdsClusterEndpoint#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/rds_cluster_endpoint#tags_all RdsClusterEndpoint#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

