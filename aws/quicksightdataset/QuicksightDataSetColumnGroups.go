package quicksightdataset


type QuicksightDataSetColumnGroups struct {
	// geo_spatial_column_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_set#geo_spatial_column_group QuicksightDataSet#geo_spatial_column_group}
	GeoSpatialColumnGroup *QuicksightDataSetColumnGroupsGeoSpatialColumnGroup `field:"optional" json:"geoSpatialColumnGroup" yaml:"geoSpatialColumnGroup"`
}

