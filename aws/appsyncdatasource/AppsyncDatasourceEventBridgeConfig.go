package appsyncdatasource


type AppsyncDatasourceEventBridgeConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#event_bus_arn AppsyncDatasource#event_bus_arn}.
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
}

