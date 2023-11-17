// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksapplication


type OpsworksApplicationAppSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/opsworks_application#type OpsworksApplication#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/opsworks_application#password OpsworksApplication#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/opsworks_application#revision OpsworksApplication#revision}.
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/opsworks_application#ssh_key OpsworksApplication#ssh_key}.
	SshKey *string `field:"optional" json:"sshKey" yaml:"sshKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/opsworks_application#url OpsworksApplication#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/opsworks_application#username OpsworksApplication#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

