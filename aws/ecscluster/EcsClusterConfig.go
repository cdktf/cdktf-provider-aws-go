// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecscluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/ecs_cluster#name EcsCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/ecs_cluster#configuration EcsCluster#configuration}
	Configuration *EcsClusterConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/ecs_cluster#id EcsCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// service_connect_defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/ecs_cluster#service_connect_defaults EcsCluster#service_connect_defaults}
	ServiceConnectDefaults *EcsClusterServiceConnectDefaults `field:"optional" json:"serviceConnectDefaults" yaml:"serviceConnectDefaults"`
	// setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/ecs_cluster#setting EcsCluster#setting}
	Setting interface{} `field:"optional" json:"setting" yaml:"setting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/ecs_cluster#tags EcsCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/ecs_cluster#tags_all EcsCluster#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

