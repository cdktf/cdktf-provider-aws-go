// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferserver


type TransferServerS3StorageOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/transfer_server#directory_listing_optimization TransferServer#directory_listing_optimization}.
	DirectoryListingOptimization *string `field:"optional" json:"directoryListingOptimization" yaml:"directoryListingOptimization"`
}

