// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsappmeshgatewayroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsAppmeshGatewayRouteConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/appmesh_gateway_route#mesh_name DataAwsAppmeshGatewayRoute#mesh_name}.
	MeshName *string `field:"required" json:"meshName" yaml:"meshName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/appmesh_gateway_route#name DataAwsAppmeshGatewayRoute#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/appmesh_gateway_route#virtual_gateway_name DataAwsAppmeshGatewayRoute#virtual_gateway_name}.
	VirtualGatewayName *string `field:"required" json:"virtualGatewayName" yaml:"virtualGatewayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/appmesh_gateway_route#id DataAwsAppmeshGatewayRoute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/appmesh_gateway_route#mesh_owner DataAwsAppmeshGatewayRoute#mesh_owner}.
	MeshOwner *string `field:"optional" json:"meshOwner" yaml:"meshOwner"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/appmesh_gateway_route#region DataAwsAppmeshGatewayRoute#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/appmesh_gateway_route#tags DataAwsAppmeshGatewayRoute#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

