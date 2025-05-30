// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetRefreshPropertiesRefreshConfiguration struct {
	// incremental_refresh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/quicksight_data_set#incremental_refresh QuicksightDataSet#incremental_refresh}
	IncrementalRefresh *QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefresh `field:"required" json:"incrementalRefresh" yaml:"incrementalRefresh"`
}

