// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcendpointserviceallowedprincipal

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcEndpointServiceAllowedPrincipalConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/vpc_endpoint_service_allowed_principal#principal_arn VpcEndpointServiceAllowedPrincipal#principal_arn}.
	PrincipalArn *string `field:"required" json:"principalArn" yaml:"principalArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/vpc_endpoint_service_allowed_principal#vpc_endpoint_service_id VpcEndpointServiceAllowedPrincipal#vpc_endpoint_service_id}.
	VpcEndpointServiceId *string `field:"required" json:"vpcEndpointServiceId" yaml:"vpcEndpointServiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/vpc_endpoint_service_allowed_principal#id VpcEndpointServiceAllowedPrincipal#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

