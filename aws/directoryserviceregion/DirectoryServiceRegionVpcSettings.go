// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryserviceregion


type DirectoryServiceRegionVpcSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/directory_service_region#subnet_ids DirectoryServiceRegion#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/directory_service_region#vpc_id DirectoryServiceRegion#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

