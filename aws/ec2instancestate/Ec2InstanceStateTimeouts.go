// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2instancestate


type Ec2InstanceStateTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ec2_instance_state#create Ec2InstanceState#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ec2_instance_state#delete Ec2InstanceState#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ec2_instance_state#update Ec2InstanceState#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

