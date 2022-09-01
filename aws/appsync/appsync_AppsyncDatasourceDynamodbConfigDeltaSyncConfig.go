package appsync


type AppsyncDatasourceDynamodbConfigDeltaSyncConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#delta_sync_table_name AppsyncDatasource#delta_sync_table_name}.
	DeltaSyncTableName *string `field:"required" json:"deltaSyncTableName" yaml:"deltaSyncTableName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#base_table_ttl AppsyncDatasource#base_table_ttl}.
	BaseTableTtl *float64 `field:"optional" json:"baseTableTtl" yaml:"baseTableTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#delta_sync_table_ttl AppsyncDatasource#delta_sync_table_ttl}.
	DeltaSyncTableTtl *float64 `field:"optional" json:"deltaSyncTableTtl" yaml:"deltaSyncTableTtl"`
}

