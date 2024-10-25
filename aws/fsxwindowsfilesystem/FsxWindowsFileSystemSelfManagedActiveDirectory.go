// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxwindowsfilesystem


type FsxWindowsFileSystemSelfManagedActiveDirectory struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/fsx_windows_file_system#dns_ips FsxWindowsFileSystem#dns_ips}.
	DnsIps *[]*string `field:"required" json:"dnsIps" yaml:"dnsIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/fsx_windows_file_system#domain_name FsxWindowsFileSystem#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/fsx_windows_file_system#password FsxWindowsFileSystem#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/fsx_windows_file_system#username FsxWindowsFileSystem#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/fsx_windows_file_system#file_system_administrators_group FsxWindowsFileSystem#file_system_administrators_group}.
	FileSystemAdministratorsGroup *string `field:"optional" json:"fileSystemAdministratorsGroup" yaml:"fileSystemAdministratorsGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/fsx_windows_file_system#organizational_unit_distinguished_name FsxWindowsFileSystem#organizational_unit_distinguished_name}.
	OrganizationalUnitDistinguishedName *string `field:"optional" json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
}

