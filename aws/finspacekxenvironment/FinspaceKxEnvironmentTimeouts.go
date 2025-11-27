// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package finspacekxenvironment


type FinspaceKxEnvironmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/finspace_kx_environment#create FinspaceKxEnvironment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/finspace_kx_environment#delete FinspaceKxEnvironment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/finspace_kx_environment#update FinspaceKxEnvironment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

