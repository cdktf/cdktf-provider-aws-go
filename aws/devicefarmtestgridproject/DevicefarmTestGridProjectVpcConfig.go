// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devicefarmtestgridproject


type DevicefarmTestGridProjectVpcConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/devicefarm_test_grid_project#security_group_ids DevicefarmTestGridProject#security_group_ids}.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/devicefarm_test_grid_project#subnet_ids DevicefarmTestGridProject#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/devicefarm_test_grid_project#vpc_id DevicefarmTestGridProject#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

