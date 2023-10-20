// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryservicedirectory


type DirectoryServiceDirectoryConnectSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/directory_service_directory#customer_dns_ips DirectoryServiceDirectory#customer_dns_ips}.
	CustomerDnsIps *[]*string `field:"required" json:"customerDnsIps" yaml:"customerDnsIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/directory_service_directory#customer_username DirectoryServiceDirectory#customer_username}.
	CustomerUsername *string `field:"required" json:"customerUsername" yaml:"customerUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/directory_service_directory#subnet_ids DirectoryServiceDirectory#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/directory_service_directory#vpc_id DirectoryServiceDirectory#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

