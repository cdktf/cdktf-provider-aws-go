// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eipdomainname

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EipDomainNameConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/eip_domain_name#allocation_id EipDomainName#allocation_id}.
	AllocationId *string `field:"required" json:"allocationId" yaml:"allocationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/eip_domain_name#domain_name EipDomainName#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/eip_domain_name#timeouts EipDomainName#timeouts}
	Timeouts *EipDomainNameTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

