// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceParametersAwsIotAnalytics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/quicksight_data_source#data_set_name QuicksightDataSource#data_set_name}.
	DataSetName *string `field:"required" json:"dataSetName" yaml:"dataSetName"`
}

