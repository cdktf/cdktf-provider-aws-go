// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxontapstoragevirtualmachine


type FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/fsx_ontap_storage_virtual_machine#dns_ips FsxOntapStorageVirtualMachine#dns_ips}.
	DnsIps *[]*string `field:"required" json:"dnsIps" yaml:"dnsIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/fsx_ontap_storage_virtual_machine#domain_name FsxOntapStorageVirtualMachine#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/fsx_ontap_storage_virtual_machine#password FsxOntapStorageVirtualMachine#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/fsx_ontap_storage_virtual_machine#username FsxOntapStorageVirtualMachine#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/fsx_ontap_storage_virtual_machine#file_system_administrators_group FsxOntapStorageVirtualMachine#file_system_administrators_group}.
	FileSystemAdministratorsGroup *string `field:"optional" json:"fileSystemAdministratorsGroup" yaml:"fileSystemAdministratorsGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/fsx_ontap_storage_virtual_machine#organizational_unit_distinguished_name FsxOntapStorageVirtualMachine#organizational_unit_distinguished_name}.
	OrganizationalUnitDistinguishedName *string `field:"optional" json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
}

