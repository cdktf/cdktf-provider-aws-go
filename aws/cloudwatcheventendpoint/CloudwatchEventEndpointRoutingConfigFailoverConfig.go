package cloudwatcheventendpoint


type CloudwatchEventEndpointRoutingConfigFailoverConfig struct {
	// primary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/cloudwatch_event_endpoint#primary CloudwatchEventEndpoint#primary}
	Primary *CloudwatchEventEndpointRoutingConfigFailoverConfigPrimary `field:"required" json:"primary" yaml:"primary"`
	// secondary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/cloudwatch_event_endpoint#secondary CloudwatchEventEndpoint#secondary}
	Secondary *CloudwatchEventEndpointRoutingConfigFailoverConfigSecondary `field:"required" json:"secondary" yaml:"secondary"`
}

