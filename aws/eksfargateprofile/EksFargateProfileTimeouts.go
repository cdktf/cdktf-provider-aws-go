// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eksfargateprofile


type EksFargateProfileTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/eks_fargate_profile#create EksFargateProfile#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/eks_fargate_profile#delete EksFargateProfile#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

