// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmscustomkeystore


type KmsCustomKeyStoreXksProxyAuthenticationCredential struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/kms_custom_key_store#access_key_id KmsCustomKeyStore#access_key_id}.
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/kms_custom_key_store#raw_secret_access_key KmsCustomKeyStore#raw_secret_access_key}.
	RawSecretAccessKey *string `field:"required" json:"rawSecretAccessKey" yaml:"rawSecretAccessKey"`
}

