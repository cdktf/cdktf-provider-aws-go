// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsconnectcontactflowmodule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsConnectContactFlowModuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/data-sources/connect_contact_flow_module#instance_id DataAwsConnectContactFlowModule#instance_id}.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/data-sources/connect_contact_flow_module#contact_flow_module_id DataAwsConnectContactFlowModule#contact_flow_module_id}.
	ContactFlowModuleId *string `field:"optional" json:"contactFlowModuleId" yaml:"contactFlowModuleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/data-sources/connect_contact_flow_module#id DataAwsConnectContactFlowModule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/data-sources/connect_contact_flow_module#name DataAwsConnectContactFlowModule#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/data-sources/connect_contact_flow_module#tags DataAwsConnectContactFlowModule#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

