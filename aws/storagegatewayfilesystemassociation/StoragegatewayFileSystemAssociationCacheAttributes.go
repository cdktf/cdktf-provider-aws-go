// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagegatewayfilesystemassociation


type StoragegatewayFileSystemAssociationCacheAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/storagegateway_file_system_association#cache_stale_timeout_in_seconds StoragegatewayFileSystemAssociation#cache_stale_timeout_in_seconds}.
	CacheStaleTimeoutInSeconds *float64 `field:"optional" json:"cacheStaleTimeoutInSeconds" yaml:"cacheStaleTimeoutInSeconds"`
}

