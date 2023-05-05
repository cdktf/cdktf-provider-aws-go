package fsxfilecache


type FsxFileCacheDataRepositoryAssociationNfs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/fsx_file_cache#version FsxFileCache#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/fsx_file_cache#dns_ips FsxFileCache#dns_ips}.
	DnsIps *[]*string `field:"optional" json:"dnsIps" yaml:"dnsIps"`
}

