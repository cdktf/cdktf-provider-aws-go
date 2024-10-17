// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawssecretsmanagersecrets


type DataAwsSecretsmanagerSecretsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/data-sources/secretsmanager_secrets#name DataAwsSecretsmanagerSecrets#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/data-sources/secretsmanager_secrets#values DataAwsSecretsmanagerSecrets#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

