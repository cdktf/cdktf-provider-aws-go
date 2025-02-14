// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasynclocationefs


type DatasyncLocationEfsEc2Config struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/datasync_location_efs#security_group_arns DatasyncLocationEfs#security_group_arns}.
	SecurityGroupArns *[]*string `field:"required" json:"securityGroupArns" yaml:"securityGroupArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/datasync_location_efs#subnet_arn DatasyncLocationEfs#subnet_arn}.
	SubnetArn *string `field:"required" json:"subnetArn" yaml:"subnetArn"`
}

