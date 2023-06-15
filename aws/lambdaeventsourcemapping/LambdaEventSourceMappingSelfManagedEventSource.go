package lambdaeventsourcemapping


type LambdaEventSourceMappingSelfManagedEventSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/lambda_event_source_mapping#endpoints LambdaEventSourceMapping#endpoints}.
	Endpoints *map[string]*string `field:"required" json:"endpoints" yaml:"endpoints"`
}

