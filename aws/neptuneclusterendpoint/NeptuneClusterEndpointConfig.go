// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package neptuneclusterendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NeptuneClusterEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/neptune_cluster_endpoint#cluster_endpoint_identifier NeptuneClusterEndpoint#cluster_endpoint_identifier}.
	ClusterEndpointIdentifier *string `field:"required" json:"clusterEndpointIdentifier" yaml:"clusterEndpointIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/neptune_cluster_endpoint#cluster_identifier NeptuneClusterEndpoint#cluster_identifier}.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/neptune_cluster_endpoint#endpoint_type NeptuneClusterEndpoint#endpoint_type}.
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/neptune_cluster_endpoint#excluded_members NeptuneClusterEndpoint#excluded_members}.
	ExcludedMembers *[]*string `field:"optional" json:"excludedMembers" yaml:"excludedMembers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/neptune_cluster_endpoint#id NeptuneClusterEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/neptune_cluster_endpoint#static_members NeptuneClusterEndpoint#static_members}.
	StaticMembers *[]*string `field:"optional" json:"staticMembers" yaml:"staticMembers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/neptune_cluster_endpoint#tags NeptuneClusterEndpoint#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/neptune_cluster_endpoint#tags_all NeptuneClusterEndpoint#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

