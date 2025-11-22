// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apprunnerservice


type ApprunnerServiceEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/apprunner_service#kms_key ApprunnerService#kms_key}.
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
}

