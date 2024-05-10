// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package glaciervault


type GlacierVaultNotification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/glacier_vault#events GlacierVault#events}.
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/glacier_vault#sns_topic GlacierVault#sns_topic}.
	SnsTopic *string `field:"required" json:"snsTopic" yaml:"snsTopic"`
}

