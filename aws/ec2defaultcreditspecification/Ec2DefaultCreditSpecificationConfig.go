// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2defaultcreditspecification

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2DefaultCreditSpecificationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/ec2_default_credit_specification#cpu_credits Ec2DefaultCreditSpecification#cpu_credits}.
	CpuCredits *string `field:"required" json:"cpuCredits" yaml:"cpuCredits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/ec2_default_credit_specification#instance_family Ec2DefaultCreditSpecification#instance_family}.
	InstanceFamily *string `field:"required" json:"instanceFamily" yaml:"instanceFamily"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/ec2_default_credit_specification#region Ec2DefaultCreditSpecification#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/ec2_default_credit_specification#timeouts Ec2DefaultCreditSpecification#timeouts}
	Timeouts *Ec2DefaultCreditSpecificationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

