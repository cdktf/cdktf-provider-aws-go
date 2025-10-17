// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceVpcConnectionProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/quicksight_data_source#vpc_connection_arn QuicksightDataSource#vpc_connection_arn}.
	VpcConnectionArn *string `field:"required" json:"vpcConnectionArn" yaml:"vpcConnectionArn"`
}

