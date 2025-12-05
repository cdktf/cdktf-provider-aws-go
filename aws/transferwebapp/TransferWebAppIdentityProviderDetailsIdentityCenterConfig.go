// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferwebapp


type TransferWebAppIdentityProviderDetailsIdentityCenterConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/transfer_web_app#instance_arn TransferWebApp#instance_arn}.
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/transfer_web_app#role TransferWebApp#role}.
	Role *string `field:"optional" json:"role" yaml:"role"`
}

