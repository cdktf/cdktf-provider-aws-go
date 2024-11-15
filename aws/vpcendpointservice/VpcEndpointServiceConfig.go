// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcendpointservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcEndpointServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/vpc_endpoint_service#acceptance_required VpcEndpointService#acceptance_required}.
	AcceptanceRequired interface{} `field:"required" json:"acceptanceRequired" yaml:"acceptanceRequired"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/vpc_endpoint_service#allowed_principals VpcEndpointService#allowed_principals}.
	AllowedPrincipals *[]*string `field:"optional" json:"allowedPrincipals" yaml:"allowedPrincipals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/vpc_endpoint_service#gateway_load_balancer_arns VpcEndpointService#gateway_load_balancer_arns}.
	GatewayLoadBalancerArns *[]*string `field:"optional" json:"gatewayLoadBalancerArns" yaml:"gatewayLoadBalancerArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/vpc_endpoint_service#id VpcEndpointService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/vpc_endpoint_service#network_load_balancer_arns VpcEndpointService#network_load_balancer_arns}.
	NetworkLoadBalancerArns *[]*string `field:"optional" json:"networkLoadBalancerArns" yaml:"networkLoadBalancerArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/vpc_endpoint_service#private_dns_name VpcEndpointService#private_dns_name}.
	PrivateDnsName *string `field:"optional" json:"privateDnsName" yaml:"privateDnsName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/vpc_endpoint_service#supported_ip_address_types VpcEndpointService#supported_ip_address_types}.
	SupportedIpAddressTypes *[]*string `field:"optional" json:"supportedIpAddressTypes" yaml:"supportedIpAddressTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/vpc_endpoint_service#tags VpcEndpointService#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/vpc_endpoint_service#tags_all VpcEndpointService#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/vpc_endpoint_service#timeouts VpcEndpointService#timeouts}
	Timeouts *VpcEndpointServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

