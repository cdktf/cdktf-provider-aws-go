// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferuser


type TransferUserTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/transfer_user#delete TransferUser#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

