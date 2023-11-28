// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package finspacekxcluster


type FinspaceKxClusterVpcConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/finspace_kx_cluster#ip_address_type FinspaceKxCluster#ip_address_type}.
	IpAddressType *string `field:"required" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/finspace_kx_cluster#security_group_ids FinspaceKxCluster#security_group_ids}.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/finspace_kx_cluster#subnet_ids FinspaceKxCluster#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/finspace_kx_cluster#vpc_id FinspaceKxCluster#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

