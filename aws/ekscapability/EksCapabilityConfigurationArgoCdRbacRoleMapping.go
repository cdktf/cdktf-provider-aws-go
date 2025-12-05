// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ekscapability


type EksCapabilityConfigurationArgoCdRbacRoleMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/eks_capability#role EksCapability#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/eks_capability#identity EksCapability#identity}
	Identity interface{} `field:"optional" json:"identity" yaml:"identity"`
}

