// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferwebapp


type TransferWebAppIdentityProviderDetails struct {
	// identity_center_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/transfer_web_app#identity_center_config TransferWebApp#identity_center_config}
	IdentityCenterConfig interface{} `field:"optional" json:"identityCenterConfig" yaml:"identityCenterConfig"`
}

