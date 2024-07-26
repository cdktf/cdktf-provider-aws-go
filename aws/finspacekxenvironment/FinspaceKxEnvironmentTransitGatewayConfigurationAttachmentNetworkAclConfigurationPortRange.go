// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package finspacekxenvironment


type FinspaceKxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationPortRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/finspace_kx_environment#from FinspaceKxEnvironment#from}.
	From *float64 `field:"required" json:"from" yaml:"from"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/finspace_kx_environment#to FinspaceKxEnvironment#to}.
	To *float64 `field:"required" json:"to" yaml:"to"`
}

