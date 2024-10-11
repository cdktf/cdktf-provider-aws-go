// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryservicedirectory


type DirectoryServiceDirectoryVpcSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/directory_service_directory#subnet_ids DirectoryServiceDirectory#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/directory_service_directory#vpc_id DirectoryServiceDirectory#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

