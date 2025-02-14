// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerlistenerpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadBalancerListenerPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/load_balancer_listener_policy#load_balancer_name LoadBalancerListenerPolicy#load_balancer_name}.
	LoadBalancerName *string `field:"required" json:"loadBalancerName" yaml:"loadBalancerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/load_balancer_listener_policy#load_balancer_port LoadBalancerListenerPolicy#load_balancer_port}.
	LoadBalancerPort *float64 `field:"required" json:"loadBalancerPort" yaml:"loadBalancerPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/load_balancer_listener_policy#id LoadBalancerListenerPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/load_balancer_listener_policy#policy_names LoadBalancerListenerPolicy#policy_names}.
	PolicyNames *[]*string `field:"optional" json:"policyNames" yaml:"policyNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/load_balancer_listener_policy#triggers LoadBalancerListenerPolicy#triggers}.
	Triggers *map[string]*string `field:"optional" json:"triggers" yaml:"triggers"`
}

