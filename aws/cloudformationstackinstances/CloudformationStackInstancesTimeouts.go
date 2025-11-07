// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudformationstackinstances


type CloudformationStackInstancesTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/cloudformation_stack_instances#create CloudformationStackInstances#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/cloudformation_stack_instances#delete CloudformationStackInstances#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/cloudformation_stack_instances#update CloudformationStackInstances#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

