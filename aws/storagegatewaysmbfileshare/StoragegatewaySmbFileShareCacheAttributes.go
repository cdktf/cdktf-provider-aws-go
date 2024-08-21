// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagegatewaysmbfileshare


type StoragegatewaySmbFileShareCacheAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/storagegateway_smb_file_share#cache_stale_timeout_in_seconds StoragegatewaySmbFileShare#cache_stale_timeout_in_seconds}.
	CacheStaleTimeoutInSeconds *float64 `field:"optional" json:"cacheStaleTimeoutInSeconds" yaml:"cacheStaleTimeoutInSeconds"`
}

