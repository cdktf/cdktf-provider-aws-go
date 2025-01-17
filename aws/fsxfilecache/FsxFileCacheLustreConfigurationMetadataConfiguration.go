// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxfilecache


type FsxFileCacheLustreConfigurationMetadataConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/fsx_file_cache#storage_capacity FsxFileCache#storage_capacity}.
	StorageCapacity *float64 `field:"required" json:"storageCapacity" yaml:"storageCapacity"`
}

