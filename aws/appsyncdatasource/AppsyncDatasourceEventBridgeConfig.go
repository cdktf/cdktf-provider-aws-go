package appsyncdatasource


type AppsyncDatasourceEventBridgeConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/appsync_datasource#event_bus_arn AppsyncDatasource#event_bus_arn}.
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
}

