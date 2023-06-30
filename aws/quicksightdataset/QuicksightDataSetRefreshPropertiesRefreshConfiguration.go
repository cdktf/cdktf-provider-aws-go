package quicksightdataset


type QuicksightDataSetRefreshPropertiesRefreshConfiguration struct {
	// incremental_refresh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/quicksight_data_set#incremental_refresh QuicksightDataSet#incremental_refresh}
	IncrementalRefresh *QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefresh `field:"required" json:"incrementalRefresh" yaml:"incrementalRefresh"`
}

