// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmskey


type KmsKeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kms_key#create KmsKey#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

