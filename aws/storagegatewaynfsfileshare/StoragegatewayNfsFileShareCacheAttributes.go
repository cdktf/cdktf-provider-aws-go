package storagegatewaynfsfileshare


type StoragegatewayNfsFileShareCacheAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/storagegateway_nfs_file_share#cache_stale_timeout_in_seconds StoragegatewayNfsFileShare#cache_stale_timeout_in_seconds}.
	CacheStaleTimeoutInSeconds *float64 `field:"optional" json:"cacheStaleTimeoutInSeconds" yaml:"cacheStaleTimeoutInSeconds"`
}

