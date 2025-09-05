// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcendpointserviceprivatednsverification

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcEndpointServicePrivateDnsVerificationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/vpc_endpoint_service_private_dns_verification#service_id VpcEndpointServicePrivateDnsVerification#service_id}.
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/vpc_endpoint_service_private_dns_verification#region VpcEndpointServicePrivateDnsVerification#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/vpc_endpoint_service_private_dns_verification#timeouts VpcEndpointServicePrivateDnsVerification#timeouts}
	Timeouts *VpcEndpointServicePrivateDnsVerificationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/vpc_endpoint_service_private_dns_verification#wait_for_verification VpcEndpointServicePrivateDnsVerification#wait_for_verification}.
	WaitForVerification interface{} `field:"optional" json:"waitForVerification" yaml:"waitForVerification"`
}

