// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceSslProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/quicksight_data_source#disable_ssl QuicksightDataSource#disable_ssl}.
	DisableSsl interface{} `field:"required" json:"disableSsl" yaml:"disableSsl"`
}

