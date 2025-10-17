// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferconnector


type TransferConnectorSftpConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/transfer_connector#trusted_host_keys TransferConnector#trusted_host_keys}.
	TrustedHostKeys *[]*string `field:"optional" json:"trustedHostKeys" yaml:"trustedHostKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/transfer_connector#user_secret_id TransferConnector#user_secret_id}.
	UserSecretId *string `field:"optional" json:"userSecretId" yaml:"userSecretId"`
}

