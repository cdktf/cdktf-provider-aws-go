package cloudwatcheventendpoint


type CloudwatchEventEndpointEventBus struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/cloudwatch_event_endpoint#event_bus_arn CloudwatchEventEndpoint#event_bus_arn}.
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
}

