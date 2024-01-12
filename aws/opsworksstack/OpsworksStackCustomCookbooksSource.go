// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksstack


type OpsworksStackCustomCookbooksSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_stack#type OpsworksStack#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_stack#url OpsworksStack#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_stack#password OpsworksStack#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_stack#revision OpsworksStack#revision}.
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_stack#ssh_key OpsworksStack#ssh_key}.
	SshKey *string `field:"optional" json:"sshKey" yaml:"sshKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_stack#username OpsworksStack#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

