// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcipam


type VpcIpamOperatingRegions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/vpc_ipam#region_name VpcIpam#region_name}.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
}

