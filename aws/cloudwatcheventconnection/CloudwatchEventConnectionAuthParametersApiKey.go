package cloudwatcheventconnection


type CloudwatchEventConnectionAuthParametersApiKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/cloudwatch_event_connection#key CloudwatchEventConnection#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/cloudwatch_event_connection#value CloudwatchEventConnection#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

